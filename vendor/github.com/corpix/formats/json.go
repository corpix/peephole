package formats

import (
	"encoding/json"
)

// JSONFormat is a JSON marshaler.
type JSONFormat uint8

// Marshal serializes data represented by v into slice of bytes.
func (j *JSONFormat) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// Unmarshal deserializes data represented by data byte slice into v.
func (j *JSONFormat) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

// Name returns a format name which is used to identify this format
// in the package.
func (j *JSONFormat) Name() string {
	return JSON
}

// NewJSON constructs a new JSON format marshaler.
func NewJSON() *JSONFormat { return new(JSONFormat) }
