package socks

import (
	"bufio"
	"fmt"
	"net"

	"github.com/corpix/loggers"
	"github.com/corpix/loggers/logger/prefixwrapper"
	"github.com/corpix/reflect"
)

const (
	socks5Version = uint8(5)
)

// Server is reponsible for accepting connections and handling
// the details of the SOCKS5 protocol.
type Server struct {
	Params Params

	log         loggers.Logger
	authMethods map[uint8]Authenticator
}

// New creates a new Server.
func New(p Params) *Server {
	var (
		ps = ParamsWithDefaults(p)
		s  = &Server{
			Params: ps,
			log:    prefixwrapper.New("Server: ", ps.Logger),
			authMethods: make(
				map[uint8]Authenticator,
				len(ps.Authenticators),
			),
		}
	)

	for _, v := range ps.Authenticators {
		s.authMethods[v.GetCode()] = v
	}

	return s
}

// ListenAndServe is used to create a listener and serve on it.
func (s *Server) ListenAndServe(network, addr string) error {
	l, err := net.Listen(network, addr)
	if err != nil {
		return err
	}
	return s.Serve(l)
}

// Serve is used to serve connections from a listener.
func (s *Server) Serve(l net.Listener) error {
	for {
		conn, err := l.Accept()
		if err != nil {
			return err
		}

		// FIXME: Pool
		go s.serveConnection(conn)
	}

	return nil
}

// serveConnection serves a connection and logs errors.
func (s *Server) serveConnection(conn net.Conn) {
	defer func() {
		if r := recover(); r != nil {
			s.log.Error(
				NewErrServingConnection(
					conn,
					fmt.Errorf("%s", r),
				),
			)
		}
	}()

	err := s.ServeConnection(conn)
	if err != nil {
		s.log.Error(NewErrServingConnection(conn, err))
	}
}

// ServeConnection is used to serve a single connection.
func (s *Server) ServeConnection(conn net.Conn) error {
	defer conn.Close()

	var (
		bufConn = bufio.NewReader(conn)
		version = []byte{0}
	)

	if _, err := bufConn.Read(version); err != nil {
		return NewErrReadingProtocolVersion(err)
	}

	if version[0] != socks5Version {
		return NewErrUnsupportedSocksVersion(socks5Version, version[0])
	}

	authContext, err := s.authenticate(conn, bufConn)
	if err != nil {
		return NewErrAuthenticationFailed(err)
	}

	request, err := NewRequest(bufConn)
	if err != nil {
		if err == unrecognizedAddrType {
			if err := sendReply(conn, addrTypeNotSupported, nil); err != nil {
				return NewErrReplyFailed(err)
			}
		}
		return NewErrRequestFailed(err)
	}

	request.AuthContext = authContext
	remoteAddr := conn.RemoteAddr()

	switch client := remoteAddr.(type) {
	case *net.TCPAddr:
		request.RemoteAddr = &AddrSpec{IP: client.IP, Port: client.Port}
	default:
		return reflect.NewErrCanNotAssertType(
			remoteAddr,
			reflect.TypeOf(&net.TCPAddr{}),
		)
	}

	err = s.handleRequest(conn, request)
	if err != nil {
		return NewErrRequestHandlerFailed(err)
	}

	return nil
}
