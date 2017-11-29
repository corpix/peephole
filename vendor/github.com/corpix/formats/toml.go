package formats

import (
	// XXX: We use this parser because
	// github.com/BurntSushi/toml API
	// is a inconsistent crap.
	"github.com/naoina/toml"
)

// TOMLFormat is a TOML marshaler.
type TOMLFormat struct{}

// Marshal serializes data represented by v into slice of bytes.
func (y *TOMLFormat) Marshal(v interface{}) ([]byte, error) {
	return toml.Marshal(v)
}

// Unmarshal deserializes data represented by data byte slice into v.
func (y *TOMLFormat) Unmarshal(data []byte, v interface{}) error {
	return toml.Unmarshal(data, v)
}

// Name returns a format name which is used to identify this format
// in the package.
func (y *TOMLFormat) Name() string {
	return TOML
}

// NewTOML constructs a new TOML format marshaler.
func NewTOML() *TOMLFormat { return &TOMLFormat{} }
