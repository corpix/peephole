package config

import (
	"bytes"
	"fmt"
	"os"

	"github.com/corpix/revip"
	yaml "gopkg.in/yaml.v2"

	"github.com/corpix/peephole/pkg/log"
	"github.com/corpix/peephole/pkg/proxy"
)

const (
	EnvironPrefix string = "PEEPHOLE"
)

var NewEncoder = yaml.NewEncoder

type Config struct {
	Log    log.Config
	Listen string
	Proxy  *proxy.Config
}

func (c *Config) Default() {
loop:
	for {
		switch {
		case c.Listen == "":
			c.Listen = "127.0.0.1:1080"
		case c.Proxy == nil:
			c.Proxy = &proxy.Config{}
		default:
			break loop
		}
	}
}

func (c *Config) Validate() error {
	if c.Listen == "" {
		return fmt.Errorf("listen address is not defined")
	}
	return nil
}

//

func Default() (*Config, error) {
	c := &Config{}
	err := revip.Postprocess(
		c,
		revip.WithDefaults(),
	)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func Load(path string) (*Config, error) {
	c := &Config{}

	fd, err := os.Open(path)
	if nil != err {
		return nil, err
	}
	defer fd.Close()

	_, err = revip.Load(
		c,
		revip.FromReader(fd, revip.YamlUnmarshaler),
		revip.FromEnviron(EnvironPrefix),
	)
	if err != nil {
		return nil, err
	}

	err = revip.Postprocess(
		c,
		revip.WithDefaults(),
		revip.WithValidation(),
	)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func Encode(c *Config) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	enc := NewEncoder(buf)
	defer enc.Close()

	err := enc.Encode(c)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func Show(c *Config) error {
	buf, err := Encode(c)
	if err != nil {
		return err
	}

	_, err = os.Stdout.Write(buf)
	return err
}
