package formats

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func TestNewFromPath(t *testing.T) {
	samples := []struct {
		path   string
		format Format
		err    error
	}{

		// json

		{
			"foo/bar.json",
			NewJSON(),
			nil,
		},
		{
			".json",
			NewJSON(),
			nil,
		},
		{
			"foo/barjson",
			nil,
			NewErrFormatNameIsEmpty(),
		},

		// yaml

		{
			"foo/bar.yaml",
			NewYAML(),
			nil,
		},
		{
			"foo/bar.yml",
			NewYAML(),
			nil,
		},
		{
			"foo/baryaml",
			nil,
			NewErrFormatNameIsEmpty(),
		},

		// other

		{
			"foo/bar.lol",
			nil,
			NewErrNotSupported("lol"),
		},
	}

	for k, sample := range samples {
		msg := spew.Sdump(k, sample)

		format, err := NewFromPath(sample.path)

		assert.EqualValues(t, sample.err, err, msg)
		assert.EqualValues(t, sample.format, format, msg)
	}
}

func TestNew(t *testing.T) {
	samples := []struct {
		name   string
		format Format
		err    error
	}{

		// json

		{
			"json",
			NewJSON(),
			nil,
		},
		{
			"JSON",
			nil,
			NewErrNotSupported("JSON"),
		},
		{
			"",
			nil,
			NewErrFormatNameIsEmpty(),
		},

		// yaml
		{
			"yaml",
			NewYAML(),
			nil,
		},
		{
			"yml",
			nil,
			NewErrNotSupported("yml"),
		},
		{
			"YAML",
			nil,
			NewErrNotSupported("YAML"),
		},
		{
			"",
			nil,
			NewErrFormatNameIsEmpty(),
		},
	}

	for k, sample := range samples {
		msg := spew.Sdump(k, sample)

		format, err := New(sample.name)

		assert.EqualValues(t, sample.err, err, msg)
		assert.EqualValues(t, sample.format, format, msg)
	}
}
