package nsq

import (
	"fmt"
	"log"
	"os"

	"github.com/nsqio/go-nsq"

	"github.com/corpix/loggers"
	"github.com/corpix/loggers/encoder"
)

// Config configuration for logger.
type Config struct {
	Level Level
	Topic string
}

// Nsq is a logger which logs to nsq message queue.
type Nsq struct {
	config   Config
	producer *nsq.Producer
	fallback loggers.Logger
	encoder.Encoder
}

// Write slice of bytes into the logger and return number of written
// bytes and error value of present.
func (l *Nsq) Write(buf []byte) (int, error) {
	var (
		n      int
		errStr string
		err    error
	)

	n, err = len(buf), l.producer.Publish(
		l.config.Topic,
		buf,
	)
	if err != nil {
		errStr = err.Error()
		switch l.fallback {
		case nil:
			log.Print(errStr)
			log.Print(string(buf))
		default:
			l.fallback.Error(err)
			n, err = l.fallback.Write(buf)
		}
	}
	return n, err
}

func (l *Nsq) Debugf(s string, xs ...interface{}) {
	l.log(DebugLevel, fmt.Sprintf(s, xs))
}

func (l *Nsq) Printf(s string, xs ...interface{}) {
	l.log(InfoLevel, fmt.Sprintf(s, xs))
}

func (l *Nsq) Errorf(s string, xs ...interface{}) {
	l.log(ErrorLevel, fmt.Sprintf(s, xs))
}

func (l *Nsq) Fatalf(s string, xs ...interface{}) {
	defer os.Exit(1)
	l.log(FatalLevel, fmt.Sprintf(s, xs))
}

func (l *Nsq) Debug(xs ...interface{}) {
	l.log(DebugLevel, xs)
}

func (l *Nsq) Print(xs ...interface{}) {
	l.log(InfoLevel, xs)
}

func (l *Nsq) Error(xs ...interface{}) {
	l.log(ErrorLevel, xs)
}

func (l *Nsq) Fatal(xs ...interface{}) {
	defer os.Exit(1)
	l.log(FatalLevel, xs)
}

func (l *Nsq) log(lvl Level, payload interface{}) {
	var (
		m   []byte
		err error
	)

	if lvl < l.config.Level {
		return
	}

	m, err = l.Encoder.Encode(NewMessage(lvl, payload))
	if err != nil {
		log.Print(err)
		return
	}

	_, err = l.Write(m)
	if err != nil {
		log.Print(err)
		return
	}
}

// Level returns a current logger level number.
func (l *Nsq) Level() interface{} {
	return l.config.Level
}

// New wraps nsq producer with Logger interface implementation.
func New(c Config, p *nsq.Producer, f loggers.Logger, e encoder.Encoder) loggers.Logger {
	return &Nsq{c, p, f, e}
}
