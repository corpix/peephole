package writer

import (
	"io"
)

type WriterChannel struct {
	io.Writer
	Channel chan *[]byte
}

func NewWriterChannel(w io.Writer, c chan *[]byte) *WriterChannel {
	return &WriterChannel{
		Writer:  w,
		Channel: c,
	}
}
