package reflect

import (
	"fmt"
)

type ErrWrongArgumentsQuantity struct {
	Want int
	Got  int
}

func (e *ErrWrongArgumentsQuantity) Error() string {
	return fmt.Sprintf(
		"Wrong arguments quantity, want '%d', got '%d'",
		e.Want,
		e.Got,
	)
}

func NewErrWrongArgumentsQuantity(want int, got int) *ErrWrongArgumentsQuantity {
	return &ErrWrongArgumentsQuantity{
		Want: want,
		Got:  got,
	}
}
