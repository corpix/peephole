package socks

import (
	"fmt"
)

type ErrReplyFailed struct {
	Err error
}

func (e ErrReplyFailed) Error() string {
	return fmt.Sprintf(
		"Reply failed: %s",
		e.Err,
	)
}

func NewErrReplyFailed(err error) ErrReplyFailed {
	return ErrReplyFailed{err}
}
