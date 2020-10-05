package socks

import (
	"fmt"
)

type ErrReadingAuthMethods struct {
	Err error
}

func (e ErrReadingAuthMethods) Error() string {
	return fmt.Sprintf(
		"Error reading authentication methods: %s",
		e.Err,
	)
}

func NewErrReadingAuthMethods(err error) ErrReadingAuthMethods {
	return ErrReadingAuthMethods{err}
}
