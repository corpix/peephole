package errors

import (
	"fmt"
	"os"

	"github.com/cockroachdb/errors"
)

var (
	New    = errors.New
	Errorf = errors.Errorf
	Wrap   = errors.Wrap
	Wrapf  = errors.Wrapf
	Cause  = errors.Cause
)

func Fatal(err error) {
	fmt.Fprintf(os.Stderr, "fatal error: %s\n", err)
	os.Exit(1)
}
