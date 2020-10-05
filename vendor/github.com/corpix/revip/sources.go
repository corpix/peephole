package revip

import (
	"io"
	"io/ioutil"
	"os"
	"syscall"

	json "encoding/json"
	yaml "github.com/go-yaml/yaml"
	env "github.com/kelseyhightower/envconfig"
	toml "github.com/pelletier/go-toml"
)

// Unmarshaler describes a generic unmarshal interface
// which could be used to extend supported formats by defining new `Option`
// implementations.
type Unmarshaler = func(in []byte, v interface{}) error

var (
	JsonUnmarshaler Unmarshaler = json.Unmarshal
	YamlUnmarshaler Unmarshaler = yaml.Unmarshal
	TomlUnmarshaler Unmarshaler = toml.Unmarshal
)

// FromReader is a `Source` constructor which creates a thunk
// to read configuration from `r` and decode it with `f` unmarshaler.
// Current implementation buffers all data in memory.
func FromReader(r io.Reader, f Unmarshaler) Option {
	return func(c Config, m ...OptionMeta) error {
		buf, err := ioutil.ReadAll(r)
		if err != nil {
			return err
		}

		return f(buf, c)
	}
}

// FromFile is a `Source` constructor which creates a thunk
// to read configuration from file addressable by `path` and
// decodes it with `f` unmarshaler.
func FromFile(path string, f Unmarshaler) Option {
	return func(c Config, m ...OptionMeta) error {
		r, err := os.Open(path)
		switch e := err.(type) {
		case *os.PathError:
			if e.Err == syscall.ENOENT {
				return &ErrFileNotFound{
					Path: path,
					Err:  err,
				}
			}
		case nil:
		default:
			return err
		}
		defer r.Close()

		return FromReader(r, f)(c)
	}
}

// FromEnviron is a `Source` constructor which creates a thunk
// to read configuration from environment.
// It uses `github.com/kelseyhightower/envconfig` underneath.
func FromEnviron(prefix string) Option {
	return func(c Config, m ...OptionMeta) error {
		return env.Process(prefix, c)
	}
}
