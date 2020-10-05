package socks

import (
	"fmt"
	"net"
)

type ErrServingConnection struct {
	Conn net.Conn
	Err  error
}

func (e ErrServingConnection) Error() string {
	return fmt.Sprintf(
		"Error while serving connection '%s': %s",
		e.Conn.RemoteAddr(),
		e.Err,
	)
}

func NewErrServingConnection(conn net.Conn, err error) ErrServingConnection {
	return ErrServingConnection{conn, err}
}
