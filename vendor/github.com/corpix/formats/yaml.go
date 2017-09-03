package formats

import (
	"gopkg.in/yaml.v2"
)

// YAMLFormat is a YAML marshaler.
type YAMLFormat uint8

// Marshal serializes data represented by v into slice of bytes.
func (y *YAMLFormat) Marshal(v interface{}) ([]byte, error) {
	return yaml.Marshal(v)
}

// Unmarshal deserializes data represented by data byte slice into v.
func (y *YAMLFormat) Unmarshal(data []byte, v interface{}) error {
	return yaml.Unmarshal(data, v)
}

// Name returns a format name which is used to identify this format
// in the package.
func (y *YAMLFormat) Name() string {
	return YAML
}

// NewYAML constructs a new YAML format marshaler.
func NewYAML() *YAMLFormat { return new(YAMLFormat) }
