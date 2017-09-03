package loggers

import (
	"io"
)

// Logger is a application level interface.
type Logger interface {
	io.Writer

	Debugf(string, ...interface{})
	Printf(string, ...interface{})
	Errorf(string, ...interface{})
	Fatalf(string, ...interface{})

	Debug(...interface{})
	Print(...interface{})
	Error(...interface{})
	Fatal(...interface{})
}
