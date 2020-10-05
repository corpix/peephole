package socks

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"sync"
	"time"

	"github.com/corpix/peephole/pkg/log"
)

const (
	socks5Version = uint8(5)
)

// Server is reponsible for accepting connections and handling
// the details of the SOCKS5 protocol.
type Server struct {
	Params Params

	conns       uint32
	connsLock   *sync.RWMutex
	log         log.Logger
	authMethods map[uint8]Authenticator
	done        chan struct{}
}

// New creates a new Server.
func New(p Params) *Server {
	var (
		ps = ParamsWithDefaults(p)
		s  = &Server{
			Params:    ps,
			connsLock: &sync.RWMutex{},
			log:       ps.Logger,
			authMethods: make(
				map[uint8]Authenticator,
				len(ps.Authenticators),
			),
			done: make(chan struct{}),
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
	go s.metricsWorker(s.done)

	for {
		conn, err := l.Accept()
		if err != nil {
			return err
		}

		s.log.Debug().Msgf(
			"Accepted new connection from '%s'",
			conn.RemoteAddr(),
		)

		// FIXME: Pool
		go s.serveConnection(conn)
	}

	return nil
}

// metricsWorker is a worker which emits some runtime metrics
// to statistics gathering service.
func (s *Server) metricsWorker(done <-chan struct{}) {
	for {
		select {
		case <-done:
			return
		default:
			s.connsLock.RLock()
			connectionsNumber := s.conns
			s.connsLock.RUnlock()

			s.Params.Metrics.SetGauge(
				[]string{"Server", "connections"},
				float32(connectionsNumber),
			)

			// FIXME: customizable flush interval for all metrics
			time.Sleep(1 * time.Second)
		}
	}
}

// handleError handles error with logger and reports it to statsd.
func (s *Server) handleError(err error) {
	s.log.Error().Err(err)

	switch v := err.(type) {
	case ErrServingConnection:
		err = v.Err
	}

	s.Params.Metrics.IncrCounter(
		[]string{"errors", "Server", fmt.Sprintf("%T", err)},
		1,
	)
}

// serveConnection serves a connection and logs errors.
func (s *Server) serveConnection(conn net.Conn) {
	defer s.Params.Metrics.MeasureSince(
		[]string{"Server", "ServeConnection"},
		time.Now(),
	)

	connectionNumber := uint32(0)

	s.connsLock.Lock()
	connectionNumber = s.conns + uint32(1)
	s.conns = connectionNumber
	s.connsLock.Unlock()

	s.Params.Metrics.SetGauge(
		[]string{"Server", "connections"},
		float32(connectionNumber),
	)
	defer func() {
		s.connsLock.Lock()
		lastConnectionNumber := s.conns - uint32(1)
		s.conns = lastConnectionNumber
		s.connsLock.Unlock()

		s.Params.Metrics.SetGauge(
			[]string{"Server", "connections"},
			float32(lastConnectionNumber),
		)
	}()

	defer func() {
		if r := recover(); r != nil {
			s.handleError(
				NewErrServingConnection(
					conn,
					fmt.Errorf("%s", r),
				),
			)
		}
	}()

	conn.SetReadDeadline(time.Now().Add(s.Params.ReadDeadlineDuration))
	conn.SetWriteDeadline(time.Now().Add(s.Params.WriteDeadlineDuration))

	err := s.ServeConnection(conn)
	if err != nil {
		s.handleError(NewErrServingConnection(conn, err))
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
		if err == io.EOF {
			return nil
		}
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
	client := remoteAddr.(*net.TCPAddr)
	request.RemoteAddr = &AddrSpec{IP: client.IP, Port: client.Port}

	err = s.handleRequest(conn, request)
	if err != nil {
		return NewErrRequestHandlerFailed(err)
	}

	return nil
}
