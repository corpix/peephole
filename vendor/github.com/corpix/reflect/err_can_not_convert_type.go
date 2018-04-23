package reflect

import (
	"fmt"
	"strings"
)

type ErrCanNotConvertType struct {
	value  interface{}
	from   Type
	to     Type
	reason []string
}

func (e ErrCanNotConvertType) Error() string {
	var (
		reason = strings.Join(e.reason, ", ")
	)

	if reason != "" {
		reason = ", reason: " + reason
	}

	return fmt.Sprintf(
		"Can not convert '%#v' of type '%s' to '%s'%s",
		e.value,
		e.from,
		e.to,
		reason,
	)
}

func NewErrCanNotConvertType(value interface{}, from Type, to Type, reason ...string) ErrCanNotConvertType {
	return ErrCanNotConvertType{
		value:  value,
		from:   from,
		to:     to,
		reason: reason,
	}
}
