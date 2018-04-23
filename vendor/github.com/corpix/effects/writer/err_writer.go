package writer

import (
	"fmt"
	"io"
)

type ErrWriter struct {
	Err    error
	Writer io.Writer
}

func (e ErrWriter) Error() string {
	return fmt.Sprintf(
		"Error in writer %p: %s",
		e.Writer,
		e.Err,
	)
}

func NewErrWriter(err error, w io.Writer) ErrWriter {
	return ErrWriter{
		Err:    err,
		Writer: w,
	}
}
