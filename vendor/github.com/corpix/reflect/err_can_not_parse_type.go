package reflect

import (
	"fmt"
)

type ErrCanNotParseType struct {
	t      string
	reason string
}

func (e ErrCanNotParseType) Error() string {
	return fmt.Sprintf(
		"Can not parse type '%s', reason: %s",
		e.t,
		e.reason,
	)
}

func NewErrCanNotParseType(t string, reason string) ErrCanNotParseType {
	return ErrCanNotParseType{
		t:      t,
		reason: reason,
	}
}
