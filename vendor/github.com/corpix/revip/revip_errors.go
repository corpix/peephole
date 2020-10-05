package revip

import (
	"fmt"
	"strings"
)

// ErrFileNotFound should be returned if configuration file was not found.
type ErrFileNotFound struct {
	Path string
	Err  error
}

func (e *ErrFileNotFound) Error() string {
	return fmt.Sprintf("no such file: %q", e.Path)
}

//

// ErrPathNotFound should be returned if key (path) was not found in configuration.
type ErrPathNotFound struct {
	Path string
}

func (e *ErrPathNotFound) Error() string {
	return fmt.Sprintf("no key matched for path: %q", e.Path)
}

//

// ErrPostprocess represents an error occured at the postprocess stage (set defaults, validation, etc)
type ErrPostprocess struct {
	Type string
	Path []string
	Err  error
}

func (e *ErrPostprocess) Error() string {
	return fmt.Sprintf(
		"postprocessing failed at %s: %s",
		strings.Join(append([]string{e.Type}, e.Path...), "."),
		e.Err.Error(),
	)
}
