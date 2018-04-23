package writer

import (
	"io"
	"sync"
)

// MultiWriter represents a dynamic list of Writer interfaces.
type MultiWriter struct {
	locker  *sync.RWMutex
	writers []io.Writer // FIXME: More effective solution?
}

// Add adds a Writer to the MultiWriter.
func (w *MultiWriter) Add(c io.Writer) {
	w.locker.Lock()
	defer w.locker.Unlock()

	w.writers = append(
		w.writers,
		c,
	)
}

// Remove removes Writer from the list if it exists
// and returns true in this case, otherwise it will be
// false.
func (w *MultiWriter) Remove(c io.Writer) bool {
	w.locker.Lock()
	defer w.locker.Unlock()

	for k, v := range w.writers {
		if v == c {
			if k < len(w.writers)-1 {
				w.writers = append(
					w.writers[0:k],
					w.writers[k+1:]...,
				)
			} else {
				w.writers = w.writers[0:k]
			}
			return true
		}
	}

	return false
}

// Has returns true if Writer exists in the list and false otherwise.
func (w *MultiWriter) Has(c io.Writer) bool {
	w.locker.RLock()
	defer w.locker.RUnlock()

	for _, v := range w.writers {
		if v == c {
			return true
		}
	}

	return false
}

func (w *MultiWriter) Write(buf []byte) (int, error) {
	var (
		l   = len(buf)
		n   int
		err error
	)

	w.locker.RLock()
	defer w.locker.RUnlock()

	for _, v := range w.writers {
		n, err = v.Write(buf)
		if err != nil {
			return n, err
		}

		if n != l {
			return n, io.ErrShortWrite
		}
	}

	return l, nil
}

// NewMultiWriter creates new Writer.
func NewMultiWriter(ws ...io.Writer) *MultiWriter {
	var (
		writers = []io.Writer{}
	)

	return &MultiWriter{
		locker:  &sync.RWMutex{},
		writers: append(writers, ws...),
	}
}
