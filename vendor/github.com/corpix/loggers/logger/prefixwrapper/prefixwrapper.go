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

func (w *PrefixWrapper) Write(rawMessage []byte) (int, error) {
	var (
		message = make(
			[]byte,
			len(w.bytePrefix)+len(rawMessage),
		)
	)

	copy(message, w.bytePrefix)
	copy(message[len(w.bytePrefix):], rawMessage)

	return w.log.Write(message)
}

func (w *PrefixWrapper) Debugf(rawMessage string, xs ...interface{}) {
	w.log.Debugf(w.stringPrefix+rawMessage, xs)
}

func (w *PrefixWrapper) Printf(rawMessage string, xs ...interface{}) {
	w.log.Printf(w.stringPrefix+rawMessage, xs)
}

func (w *PrefixWrapper) Errorf(rawMessage string, xs ...interface{}) {
	w.log.Errorf(w.stringPrefix+rawMessage, xs)
}

func (w *PrefixWrapper) Fatalf(rawMessage string, xs ...interface{}) {
	w.log.Fatalf(w.stringPrefix+rawMessage, xs)
}

func (w *PrefixWrapper) Debug(rawXs ...interface{}) {
	var (
		xs = make(
			[]interface{},
			len(rawXs)+1,
		)
	)

	xs[0] = w.stringPrefix
	copy(xs[1:], rawXs)

	w.log.Debug(xs...)
}

func (w *PrefixWrapper) Print(rawXs ...interface{}) {
	var (
		xs = make(
			[]interface{},
			len(rawXs)+1,
		)
	)

	xs[0] = w.stringPrefix
	copy(xs[1:], rawXs)

	w.log.Print(xs...)
}

func (w *PrefixWrapper) Error(rawXs ...interface{}) {
	var (
		xs = make(
			[]interface{},
			len(rawXs)+1,
		)
	)

	xs[0] = w.stringPrefix
	copy(xs[1:], rawXs)

	w.log.Error(xs...)
}

func (w *PrefixWrapper) Fatal(rawXs ...interface{}) {
	var (
		xs = make(
			[]interface{},
			len(rawXs)+1,
		)
	)

	xs[0] = w.stringPrefix
	copy(xs[1:], rawXs)

	w.log.Fatal(xs...)
}

// New wraps Logger with PrefixWrapper.
func New(prefix string, l loggers.Logger) loggers.Logger {
	return &PrefixWrapper{
		bytePrefix:   []byte(prefix),
		stringPrefix: prefix,
		log:          l,
	}
}
