package writer

import (
	"bytes"
	"io"
)

type TrimSuffixWriter struct {
	w      io.Writer
	suffix []byte
}

func (w *TrimSuffixWriter) Write(buf []byte) (int, error) {
	var (
		n    int
		nbuf []byte
		err  error
	)

	nbuf = bytes.TrimSuffix(buf, w.suffix)

	n, err = w.w.Write(nbuf)
	if err != nil {
		return n, err
	}

	return n + (len(buf) - len(nbuf)), nil
}

// NewTrimSuffixWriter creates new Writer.
func NewTrimSuffixWriter(w io.Writer, suffix []byte) *TrimSuffixWriter {
	return &TrimSuffixWriter{
		w:      w,
		suffix: suffix,
	}
}
