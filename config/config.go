package config

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/corpix/formats"
	"github.com/imdario/mergo"

	"github.com/corpix/go-boilerplate/logger"
)

var (
	// LoggerConfig represents default logger config.
	LoggerConfig = logger.Config{
		Level:     "info",
		Formatter: "json",
	}

	// Default represents default application config.
	Default = Config{
		Logger: LoggerConfig,
	}
)

// Config represents application configuration structure.
type Config struct {
	Logger logger.Config
}

// FromReader returns parsed config data in some `f` from reader `r`.
// It copies `Default` into the target structure before unmarshaling
// config, so it will have default values.
func FromReader(f formats.Format, r io.Reader) (Config, error) {
	var (
		c   Config
		buf []byte
		err error
	)

	buf, err = ioutil.ReadAll(r)
	if err != nil {
		return c, err
	}

	err = mergo.Merge(&c, Default)
	if err != nil {
		return c, err
	}

	err = f.Unmarshal(buf, &c)
	if err != nil {
		return c, err
	}

	return c, nil
}

// FromFile returns parsed config data from file in `path`.
func FromFile(path string) (Config, error) {
	var (
		c   Config
		f   formats.Format
		r   io.ReadWriteCloser
		err error
	)
	f, err = formats.NewFromPath(path)
	if err != nil {
		return c, err
	}

	r, err = os.Open(path)
	if err != nil {
		return c, err
	}
	defer r.Close()

	c, err = FromReader(f, r)
	if err != nil {
		return c, err
	}

	return c, nil
}
