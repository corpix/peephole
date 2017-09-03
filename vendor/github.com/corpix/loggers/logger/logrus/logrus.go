package logrus

import (
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/corpix/loggers"
)

const (
	// TextFormatter is a text formatter name for logrus logger.
	TextFormatter = "text"

	// JSONFormatter is a JSON formatter name for logrus logger.
	JSONFormatter = "json"
)

// Config is a logrus configuration which could be
// customized.
type Config struct {
	Level     string
	Formatter string
}

// Logrus is a logrus for logger that implements
// io.Writer interface.
type Logrus struct {
	*logrus.Logger
}

// Write slice of bytes into the logger and return number of written
// bytes and error value of present.
func (l *Logrus) Write(buf []byte) (int, error) {
	n := len(buf) + 1
	l.Printf("%s\n", buf)
	return n, nil
}

// Level returns a current logger level number.
func (l *Logrus) Level() interface{} {
	return l.Logger.Level
}

// New wraps logrus logger with binding.
func New(l *logrus.Logger) loggers.Logger {
	return &Logrus{l}
}

// New creates and wraps new logrus logger with binding.
func NewFromConfig(c Config) (loggers.Logger, error) {
	var (
		l   logrus.Level
		f   logrus.Formatter
		err error
	)
	l, err = logrus.ParseLevel(c.Level)
	if err != nil {
		return nil, err
	}

	switch strings.ToLower(c.Formatter) {
	case TextFormatter:
		f = &logrus.TextFormatter{}
	case JSONFormatter:
		f = &logrus.JSONFormatter{}
	case "":
		f = &logrus.TextFormatter{}
	default:
		return nil, NewErrUnknownFormatter(c.Formatter)
	}

	log := logrus.New()
	log.Level = l
	log.Formatter = f

	return New(log), nil
}
