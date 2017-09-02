package errors

import (
	"fmt"
)

type ErrNilArgument struct {
	Value interface{}
}

func (e *ErrNilArgument) Error() string {
	return fmt.Sprintf(
		"Received nil argument of type '%T'",
		e.Value,
	)
}

func NewErrNilArgument(value interface{}) error {
	return &ErrNilArgument{value}
}
