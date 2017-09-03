package logrus

import (
	"fmt"
)

type ErrUnknownFormatter struct {
	f string
}

func (e *ErrUnknownFormatter) Error() string {
	return fmt.Sprintf(
		"Unknown formatter '%s'",
		e.f,
	)
}
func NewErrUnknownFormatter(f string) error {
	return &ErrUnknownFormatter{f}
}

//
