package formats

import (
	"fmt"
)

// ErrNotSupported indicates that format with specified name is not supported.
type ErrNotSupported struct {
	Format string
}

func (e *ErrNotSupported) Error() string {
	return fmt.Sprintf(
		"Format '%s' is not supported",
		e.Format,
	)
}

// NewErrNotSupported wraps format name with ErrNotSupported.
func NewErrNotSupported(format string) error {
	return &ErrNotSupported{format}
}

//
var (
	errFormatNameIsEmpty = "Format name is empty"
)

// ErrFormatNameIsEmpty indicates that format name is an empty string.
type ErrFormatNameIsEmpty struct{}

func (e *ErrFormatNameIsEmpty) Error() string { return errFormatNameIsEmpty }

// NewErrFormatNameIsEmpty wraps format name with ErrFormatNameIsEmpty.
func NewErrFormatNameIsEmpty() error { return &ErrFormatNameIsEmpty{} }
