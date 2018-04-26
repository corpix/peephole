package socks

import (
	"fmt"
)

type ErrNoSupportedAuthMethod struct {
	Methods []Authenticator
}

func (e ErrNoSupportedAuthMethod) Error() string {
	return fmt.Sprintf(
		"No supported authentication method providen: %v",
		e.Methods,
	)
}

func NewErrNoSupportedAuthMethod(methods []Authenticator) ErrNoSupportedAuthMethod {
	return ErrNoSupportedAuthMethod{methods}
}
