package reflect

import (
	"fmt"
)

type ErrUnknownType struct {
	Value interface{}
}

func (e *ErrUnknownType) Error() string {
	return fmt.Sprintf(
		"Unknown type '%T'",
		e.Value,
	)
}

func NewErrUnknownType(v interface{}) *ErrUnknownType {
	return &ErrUnknownType{
		Value: v,
	}
}
