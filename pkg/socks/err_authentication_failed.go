package socks

import (
	"fmt"
)

type ErrAuthenticationFailed struct {
	Err error
}

func (e ErrAuthenticationFailed) Error() string {
	return fmt.Sprintf(
		"Authentication failed: %s",
		e.Err,
	)
}

func NewErrAuthenticationFailed(err error) ErrAuthenticationFailed {
	return ErrAuthenticationFailed{err}
}
