package prefixwrapper

import (
	"github.com/corpix/loggers"
)

// PrefixWrapper is a wrapper for the Logger that
// implements the same interface and adds user defined prefix
// for each call.
type PrefixWrapper struct {
	bytePrefix   []byte
	stringPrefix string
	log          loggers.Logger
}

func (l *PrefixWrapper) Write(rawMessage []byte) (int, error) {
	var (
		message = make(
			[]byte,
			len(l.bytePrefix)+len(rawMessage),
		)
	)

	copy(message, l.bytePrefix)
	copy(message[len(l.bytePrefix):], rawMessage)

	return l.log.Write(message)
}

func (l *PrefixWrapper) Debugf(rawMessage string, xs ...interface{}) {
	l.log.Debugf(l.stringPrefix+rawMessage, xs...)
}

func (l *PrefixWrapper) Printf(rawMessage string, xs ...interface{}) {
	l.log.Printf(l.stringPrefix+rawMessage, xs...)
}

func (l *PrefixWrapper) Errorf(rawMessage string, xs ...interface{}) {
	l.log.Errorf(l.stringPrefix+rawMessage, xs...)
}

func (l *PrefixWrapper) Fatalf(rawMessage string, xs ...interface{}) {
	l.log.Fatalf(l.stringPrefix+rawMessage, xs...)
}

func (l *PrefixWrapper) Debug(rawXs ...interface{}) {
	var (
		xs = make(
			[]interface{},
			len(rawXs)+1,
		)
	)

	xs[0] = l.stringPrefix
	copy(xs[1:], rawXs)

	l.log.Debug(xs...)
}

func (l *PrefixWrapper) Print(rawXs ...interface{}) {
	var (
		xs = make(
			[]interface{},
			len(rawXs)+1,
		)
	)

	xs[0] = l.stringPrefix
	copy(xs[1:], rawXs)

	l.log.Print(xs...)
}

func (l *PrefixWrapper) Error(rawXs ...interface{}) {
	var (
		xs = make(
			[]interface{},
			len(rawXs)+1,
		)
	)

	xs[0] = l.stringPrefix
	copy(xs[1:], rawXs)

	l.log.Error(xs...)
}

func (l *PrefixWrapper) Fatal(rawXs ...interface{}) {
	var (
		xs = make(
			[]interface{},
			len(rawXs)+1,
		)
	)

	xs[0] = l.stringPrefix
	copy(xs[1:], rawXs)

	l.log.Fatal(xs...)
}

// New wraps Logger with PrefixWrapper.
func New(prefix string, l loggers.Logger) loggers.Logger {
	return &PrefixWrapper{
		bytePrefix:   []byte(prefix),
		stringPrefix: prefix,
		log:          l,
	}
}
