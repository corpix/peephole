package formats

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func TestYAMLMarshal(t *testing.T) {
	samples := []struct {
		input  interface{}
		result []byte
		err    error
	}{
		{
			struct {
				Foo string `yaml:"foo"`
				Bar map[string]string
			}{
				"hello",
				map[string]string{
					"1": "one",
					"2": "two",
				},
			},
			// XXX: Lower-casing of keys is a go-yaml crap.
			// https://github.com/go-yaml/yaml/issues/148
			[]byte("foo: hello\nbar:\n  \"1\": one\n  \"2\": two\n"),
			nil,
		},
	}

	yaml := NewYAML()
	for k, sample := range samples {
		msg := spew.Sdump(k, sample)

		result, err := yaml.Marshal(sample.input)
		assert.EqualValues(t, sample.err, err, msg)
		assert.EqualValues(t, sample.result, result, msg)
	}
}
