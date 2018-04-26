package socks

import (
	"fmt"
)

type ErrRequestFailed struct {
	Err error
}

func (e ErrRequestFailed) Error() string {
	return fmt.Sprintf(
		"Request failed: %s",
		e.Err,
	)
}

func NewErrRequestFailed(err error) ErrRequestFailed {
	return ErrRequestFailed{err}
}
