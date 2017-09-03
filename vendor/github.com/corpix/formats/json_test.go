package formats

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func TestJSONMarshal(t *testing.T) {
	samples := []struct {
		input  interface{}
		result []byte
		err    error
	}{
		{
			struct {
				Foo string `json:"foo"`
				Bar map[string]string
			}{
				"hello",
				map[string]string{
					"1": "one",
					"2": "two",
				},
			},
			[]byte(`{"foo":"hello","Bar":{"1":"one","2":"two"}}`),
			nil,
		},
	}

	json := NewJSON()
	for k, sample := range samples {
		msg := spew.Sdump(k, sample)

		result, err := json.Marshal(sample.input)
		assert.EqualValues(t, sample.err, err, msg)
		assert.EqualValues(t, sample.result, result, msg)
	}
}
