package socks

import (
	"fmt"
	"net"
)

type ErrServingConnection struct {
	Err  error
	Conn net.Conn
}

func (e ErrServingConnection) Error() string {
	return fmt.Sprintf(
		"Error while serving connection: Remote addr '%s': %s",
		e.Conn.RemoteAddr(),
		e.Err,
	)
}

func NewErrServingConnection(err error, conn net.Conn) ErrServingConnection {
	return ErrServingConnection{err, conn}
}
