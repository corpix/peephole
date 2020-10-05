# revip

Dead-simple configuration loader.

It supports:

- JSON, TOML, YAML and you could add your own format unmarshaler (see `Unmarshaler` type)
- file, reader and environment sources support, also you could add your own (see `Option` type and `sources.go`)
- extendable postprocessing support (validation, defaults, see `Option` type and `postprocess.go`)
- JSON-path support

[Godoc](https://godoc.org/github.com/corpix/revip)

---

### example

```go
package main

import (
	"bytes"
	"fmt"

	"github.com/corpix/revip"
)

type (
	Foo struct {
		Bar string
		Qux bool
	}
	Config struct {
		Foo *Foo
		Baz int
		Dox []string
		Box []int
	}
)

func (c *Config) Validate() error {
	if c.Baz <= 0 {
		return fmt.Errorf("baz should be greater than zero")
	}
	return nil
}

func (c *Config) Default() {
loop:
	switch {
	case c.Foo == nil:
		c.Foo = &Foo{Bar:"bar default", Qux: true}
	default:
		break loop
	}
}

func main() {
	c := Config{
		Foo: &Foo{
			Bar: "bar",
			Qux: true,
		},
		Baz: 666,
	}

	_, err := revip.Load(
		&c,
		revip.FromReader(
			bytes.NewBuffer([]byte(`{"foo":{"qux": false}}`)),
			revip.JsonUnmarshaler,
		),
		revip.FromReader(
			bytes.NewBuffer([]byte(`box = [666,777,888]`)),
			revip.TomlUnmarshaler,
		),
		revip.FromEnviron("revip"),
	)
	if err != nil {
		panic(err)
	}

	err = revip.Postprocess(
		&c,
		revip.WithDefaults(),
		revip.WithValidation(),
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("config: %#v\n", c)
}
```

### run

```console
user@localhost> go run ./example/main.go
config: main.Config{Foo:(*main.Foo)(0xc0000c03e0), Baz:666, Dox:[]string(nil), Box:[]int{666, 777, 888}}

user@localhost> REVIP_FOO_BAR=hello go run ./example/main.go
config: main.Config{Foo:(*main.Foo)(0xc00000e440), Baz:666, Dox:[]string(nil), Box:[]int{666, 777, 888}}

user@localhost> REVIP_BOX=888,777,666 go run ./example/main.go
config: main.Config{Foo:(*main.Foo)(0xc0000c03e0), Baz:666, Dox:[]string(nil), Box:[]int{888, 777, 666}}

user@localhost> REVIP_BAZ=0 go run ./example/main.go
panic: postprocessing failed at *main.Config: baz should be greater than zero

goroutine 1 [running]:
main.main()
        /home/user/go/src/github.com/corpix/revip/example/main.go:67 +0x46a
exit status 2
```

## license

[public domain](https://unlicense.org/)
