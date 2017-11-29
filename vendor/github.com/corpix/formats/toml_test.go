package formats

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func TestTOMLMarshal(t *testing.T) {
	samples := []struct {
		input  interface{}
		result []byte
		err    error
	}{
		{
			struct {
				Foo string `toml:"foo"`
				Bar map[string]string
			}{
				"hello",
				map[string]string{
					"1": "one",
					"2": "two",
				},
			},
			[]byte("foo = \"hello\"\n\n[bar]\n1 = \"one\"\n2 = \"two\"\n"),
			nil,
		},
	}

	toml := NewTOML()
	for k, sample := range samples {
		msg := spew.Sdump(k, sample)

		result, err := toml.Marshal(sample.input)
		assert.EqualValues(t, sample.err, err, msg)
		assert.EqualValues(t, sample.result, result, msg)
	}
}
