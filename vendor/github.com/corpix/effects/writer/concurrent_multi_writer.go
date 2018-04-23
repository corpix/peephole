package writer

import (
	"context"
	"io"
	"sync"
	"time"
	"unsafe"

	"github.com/corpix/pool"
)

type writesStatus struct {
	target  int32
	current int32
	wg      *sync.WaitGroup
}

// ConcurrentMultiWriter represents a dynamic list of Writer interfaces.
type ConcurrentMultiWriter struct {
	config       ConcurrentMultiWriterConfig
	pool         *pool.Pool
	writers      []*WriterChannel
	writersLock  *sync.RWMutex
	preemptLock  *sync.RWMutex
	preempt      map[uintptr]*sync.WaitGroup
	errorHandler func(error)
}

// Add adds a Writer to the ConcurrentMultiWriter.
// Could block if it runs out of (pool.Pool.QueueSize, which could mean we have
// no free slots in pool.Pool.Workers).
func (w *ConcurrentMultiWriter) Add(wr io.Writer) {
	var (
		wc = NewWriterChannel(
			wr,
			// XXX: Closing in Remove()
			make(chan *[]byte, w.config.Backlog.Size),
		)
	)

	w.pool.Feed <- pool.NewWork(
		context.Background(),
		w.channelWriterPump(wc),
	)

	w.writersLock.Lock()
	defer w.writersLock.Unlock()
	w.writers = append(w.writers, wc)
}

// TryAdd try add a a Writer to the ConcurrentMultiWriter without blocking.
func (w *ConcurrentMultiWriter) TryAdd(wr io.Writer) bool {
	var (
		wc = NewWriterChannel(
			wr,
			// XXX: Closing in Remove()
			make(chan *[]byte, w.config.Backlog.Size),
		)
	)

	select {
	case w.pool.Feed <- pool.NewWork(
		context.Background(),
		w.channelWriterPump(wc),
	):
	default:
		return false
	}

	w.writersLock.Lock()
	defer w.writersLock.Unlock()
	w.writers = append(w.writers, wc)

	return true
}

// TryAddWithTimeout try add a a Writer to the ConcurrentMultiWriter without blocking.
func (w *ConcurrentMultiWriter) TryAddWithTimeout(wr io.Writer, timeout time.Duration) bool {
	var (
		wc = NewWriterChannel(
			wr,
			// XXX: Closing in Remove()
			make(chan *[]byte, w.config.Backlog.Size),
		)
	)

	select {
	case w.pool.Feed <- pool.NewWork(
		context.Background(),
		w.channelWriterPump(wc),
	):
	case <-time.After(timeout):
		return false
	}

	w.writersLock.Lock()
	defer w.writersLock.Unlock()
	w.writers = append(w.writers, wc)

	return true
}

// Remove removes Writer from the list if it exists
// and returns true in this case, otherwise it will be
// false.
func (w *ConcurrentMultiWriter) Remove(c io.Writer) bool {
	w.writersLock.Lock()
	defer w.writersLock.Unlock()

	for k, wr := range w.writers {
		if wr.Writer == c {
			// XXX: Created in Add()
			defer close(wr.Channel)

			if k < len(w.writers)-1 {
				w.writers = append(
					w.writers[:k],
					w.writers[k+1:]...,
				)
			} else {
				w.writers = w.writers[:k]
			}
			return true
		}
	}

	return false
}

// RemoveAll removes all writers.
func (w *ConcurrentMultiWriter) RemoveAll() bool {
	w.writersLock.Lock()
	defer w.writersLock.Unlock()

	return w.removeAll()
}

// removeAll removes all writers(not threadsafe).
func (w *ConcurrentMultiWriter) removeAll() bool {
	res := len(w.writers) > 0

	for _, wr := range w.writers {
		// XXX: Created in Add()
		defer close(wr.Channel)
	}
	w.writers = nil

	return res
}

// Has returns true if Writer exists in the list and false otherwise.
func (w *ConcurrentMultiWriter) Has(c io.Writer) bool {
	w.writersLock.RLock()
	defer w.writersLock.RUnlock()

	for _, wr := range w.writers {
		if wr.Writer == c {
			return true
		}
	}

	return false
}

func (w *ConcurrentMultiWriter) channelWriterPump(wr *WriterChannel) pool.Executor {
	return func(ctx context.Context) {
		var (
			id  uintptr
			n   int
			buf *[]byte
			err error
		)

	prelude:
		select {
		case <-ctx.Done():
			return
		case buf = <-wr.Channel:
			goto head
		}

	head:
		if buf == nil {
			return
		}

		id = uintptr(unsafe.Pointer(buf))

		n, err = wr.Writer.Write(*buf)
		if err != nil {
			w.errorHandler(NewErrWriter(err, wr.Writer))
			goto tail
		}

		if n < len(*buf) {
			w.errorHandler(NewErrWriter(io.ErrShortWrite, wr.Writer))
			goto tail
		}

	tail:
		w.preemptLock.RLock()
		w.preempt[id].Done()
		w.preemptLock.RUnlock()

		goto prelude
	}
}

func (w *ConcurrentMultiWriter) deferErrorHadler(err error) func() {
	return func() { w.errorHandler(err) }
}

func (w *ConcurrentMultiWriter) createPreemptWaitGroup(id uintptr, n int) *sync.WaitGroup {
	w.preemptLock.Lock()
	defer w.preemptLock.Unlock()

	wg, ok := w.preempt[id]
	if !ok {
		wg = &sync.WaitGroup{}
		w.preempt[id] = wg
	}
	if n > 0 {
		w.preempt[id].Add(n)
	}

	return wg
}

func (w *ConcurrentMultiWriter) removePreemptWaitGroup(id uintptr) {
	w.preemptLock.Lock()
	defer w.preemptLock.Unlock()

	delete(w.preempt, id)
}

func (w *ConcurrentMultiWriter) write(id uintptr, buf *[]byte) []func() {
	var (
		wg   *sync.WaitGroup
		errs []func()
	)

	w.writersLock.RLock()
	defer w.writersLock.RUnlock()

	if len(w.writers) == 0 {
		return nil
	}

	wg = w.createPreemptWaitGroup(id, len(w.writers))
	defer w.removePreemptWaitGroup(id)
	defer wg.Wait()

	for _, wr := range w.writers {
		select {
		case wr.Channel <- buf:
		case <-time.After(time.Duration(w.config.Backlog.AddTimeout)):
			wg.Done()
			errs = append(
				errs,
				w.deferErrorHadler(
					NewErrBacklogOverflow(w.config.Backlog.Size, wr),
				),
			)
		}
	}

	return errs
}

// Write writes a buf to each concrete writer concurrently.
// This writer is «eventually consistent», it doesn't wait until
// all data will be sent before return control.
func (w *ConcurrentMultiWriter) Write(buf []byte) (int, error) {
	var (
		id = uintptr(unsafe.Pointer(&buf))
	)

	for _, handle := range w.write(id, &buf) {
		handle()
	}

	return len(buf), nil
}

func (w *ConcurrentMultiWriter) Close() error {
	w.pool.Close()

	w.writersLock.Lock()
	defer w.writersLock.Unlock()

	w.preemptLock.RLock()
	defer w.preemptLock.RUnlock()

	for _, wg := range w.preempt {
		wg.Wait()
	}

	w.removeAll()

	return nil
}

// NewConcurrentMultiWriter creates new Writer.
func NewConcurrentMultiWriter(c ConcurrentMultiWriterConfig, errorHandler func(error), ws ...io.Writer) *ConcurrentMultiWriter {
	var (
		w = &ConcurrentMultiWriter{
			config:       c,
			pool:         pool.NewFromConfig(c.Pool),
			writersLock:  &sync.RWMutex{},
			preemptLock:  &sync.RWMutex{},
			preempt:      map[uintptr]*sync.WaitGroup{},
			errorHandler: errorHandler,
		}
	)

	for _, wr := range ws {
		w.Add(wr)
	}

	return w
}
