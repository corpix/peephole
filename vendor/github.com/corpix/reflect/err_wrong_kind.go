package reflect

import (
	"fmt"
)

type ErrWrongKind struct {
	Want Kind
	Got  Kind
}

func (e *ErrWrongKind) Error() string {
	return fmt.Sprintf(
		"Wrong kind, want '%s', got '%s'",
		e.Want,
		e.Got,
	)
}

func NewErrWrongKind(want Kind, got Kind) *ErrWrongKind {
	return &ErrWrongKind{
		Want: want,
		Got:  got,
	}
}
