package memory

import (
	"fmt"
	"sync"

	"github.com/corpix/loggers"
)

// Memory is a logger which saves all messages in memory.
// It was implemented to write tests for loggers and debug them.
type Memory struct {
	buf  [][]byte
	lock *sync.Mutex
}

func (l *Memory) Write(msg []byte) (int, error) {
	l.appendThreadSafe(msg)
	return len(msg), nil
}

func (l *Memory) Debugf(msg string, xs ...interface{}) {
	l.appendThreadSafe(l.encodef(msg, xs...))
}

func (l *Memory) Printf(msg string, xs ...interface{}) {
	l.appendThreadSafe(l.encodef(msg, xs...))
}

func (l *Memory) Errorf(msg string, xs ...interface{}) {
	l.appendThreadSafe(l.encodef(msg, xs...))
}

func (l *Memory) Fatalf(msg string, xs ...interface{}) {
	l.appendThreadSafe(l.encodef(msg, xs...))
}

func (l *Memory) Debug(xs ...interface{}) {
	l.appendThreadSafe(l.encode(xs...))
}

func (l *Memory) Print(xs ...interface{}) {
	l.appendThreadSafe(l.encode(xs...))
}

func (l *Memory) Error(xs ...interface{}) {
	l.appendThreadSafe(l.encode(xs...))
}

func (l *Memory) Fatal(xs ...interface{}) {
	l.appendThreadSafe(l.encode(xs...))
}

func (l *Memory) encodef(msg string, xs ...interface{}) []byte {
	return []byte(fmt.Sprintf(msg, xs...))
}

func (l *Memory) encode(xs ...interface{}) []byte {
	return []byte(fmt.Sprint(xs...))
}

func (l *Memory) appendThreadSafe(msg []byte) {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.buf = append(
		l.buf,
		msg,
	)
}

func (l *Memory) GetBuffer() [][]byte {
	l.lock.Lock()
	defer l.lock.Unlock()

	var (
		bufCopy = make(
			[][]byte,
			len(l.buf),
		)
	)

	copy(bufCopy, l.buf)

	return bufCopy
}

// New creates a logger which writes all messages into in-memory buffer.
// To get messages you should assert it to concrete type and call GetBuffer().
func New() loggers.Logger {
	return &Memory{
		buf:  [][]byte{},
		lock: &sync.Mutex{},
	}
}
