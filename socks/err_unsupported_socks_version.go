package socks

import (
	"fmt"
)

type ErrUnsupportedSocksVersion struct {
	Want uint8
	Got  uint8
}

func (e ErrUnsupportedSocksVersion) Error() string {
	return fmt.Sprintf("Unsupported SOCKS version, want %d, got %d", e.Want, e.Got)
}

func NewErrUnsupportedSocksVersion(want, got uint8) ErrUnsupportedSocksVersion {
	return ErrUnsupportedSocksVersion{want, got}
}
