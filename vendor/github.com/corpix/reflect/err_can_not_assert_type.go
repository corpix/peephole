package reflect

import (
	"fmt"
	"strings"
)

type ErrCanNotAssertType struct {
	value  interface{}
	to     Type
	reason []string
}

func (e ErrCanNotAssertType) Error() string {
	var (
		reason = strings.Join(e.reason, ", ")
	)

	if reason != "" {
		reason = ", reason: " + reason
	}

	return fmt.Sprintf(
		"Can not assert '%#v' to '%s'%s",
		e.value,
		e.to,
		reason,
	)
}

func NewErrCanNotAssertType(value interface{}, to Type, reason ...string) ErrCanNotAssertType {
	return ErrCanNotAssertType{
		value:  value,
		to:     to,
		reason: reason,
	}
}
