package encoder

import (
	"encoding/json"
)

// JSON is an Encoder interface implementation for JSON.
type JSON struct{}

// Encode encodes data to JSON.
func (e *JSON) Encode(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// NewJSON creates new Encoder.
func NewJSON() Encoder { return &JSON{} }
