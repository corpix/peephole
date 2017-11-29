package formats

import (
	"path/filepath"
	"strings"
)

const (
	// JSON is a JSON format name.
	JSON = "json"

	// YAML is a YAML format name.
	YAML = "yaml"

	// TOML is a TOML format name.
	TOML = "toml"
)

var (
	// Names lists supported set of formats
	// which could be used as argument to Name() and
	// also available as constants exported from the package.
	Names = []string{
		JSON,
		YAML,
		TOML,
	}

	// synonyms represents a format name synonyms mapping
	// into concrete format name.
	synonyms = map[string]string{
		"yml": YAML,
	}
)

// Format is a iterator that provides a data from source.
type Format interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
	Name() string
}

// NewFromPath returns a new Format from file extension
// in the path argument.
// It respects the synonyms for format names,
// for example: yaml format files could have extensions
// yaml or yml.
func NewFromPath(path string) (Format, error) {
	name := strings.TrimPrefix(
		filepath.Ext(path),
		".",
	)
	if synonym, ok := synonyms[name]; ok {
		name = synonym
	}

	return New(strings.ToLower(name))
}

// New create a new format marshaler/unmarshaler from name.
func New(name string) (Format, error) {
	if name == "" {
		return nil, NewErrFormatNameIsEmpty()
	}

	switch name {
	case JSON:
		return NewJSON(), nil
	case YAML:
		return NewYAML(), nil
	case TOML:
		return NewTOML(), nil
	default:
		return nil, NewErrNotSupported(name)
	}
}
