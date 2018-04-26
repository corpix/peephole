package socks

import (
	"fmt"
)

type ErrReadingProtocolVersion struct {
	Err error
}

func (e ErrReadingProtocolVersion) Error() string {
	return fmt.Sprintf(
		"Failed to read protocol version: %s",
		e.Err,
	)
}

func NewErrReadingProtocolVersion(err error) ErrReadingProtocolVersion {
	return ErrReadingProtocolVersion{err}
}
