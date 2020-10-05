package socks

import (
	"fmt"
)

type ErrRequestHandlerFailed struct {
	Err error
}

func (e ErrRequestHandlerFailed) Error() string {
	return fmt.Sprintf(
		"Request handler failed: %s",
		e.Err,
	)
}

func NewErrRequestHandlerFailed(err error) ErrRequestHandlerFailed {
	return ErrRequestHandlerFailed{err}
}
