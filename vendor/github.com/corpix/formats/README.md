formats
----------

[![Build Status](https://travis-ci.org/corpix/formats.svg?branch=master)](https://travis-ci.org/corpix/formats)

This package provides:

- command-line tool which converts stdin from format A to format B
- library with consistent API to transform data from format A to format B

Supported formats:

- `JSON`
- `YAML`
- `TOML`

## Command-line tool example

``` console
$ go get -u github.com/corpix/formats/...
...

$ echo '{"hello": ["world", {"of": "foo", "bar": true}]}' | formats --from json --to yaml
hello:
- world
- bar: true
  of: foo

$ echo '{"hello": ["world", {"of": "foo", "bar": true}]}' | formats --from json --to yaml | formats --from yaml --to json
{"hello":["world",{"bar":"true","of":"foo"}]}

$ echo -n 'hello' | formats --to hex
68656c6c6f

$ echo -n '68656c6c6f' | formats --from hex
hello

$ echo -n '68656c6c6f' | formats --from hex --to json
"hello"
```

## Library usage example

This will convert JSON to YAML:

``` go
package main

import (
	"fmt"

	"github.com/corpix/formats"
)

var (
	json = `
        {
            "name": "Danny",
            "roles": ["warrior", "worker"]
        }
    `
)

func main() {
	v := new(interface{})

	j := formats.NewJSON()
	err := j.Unmarshal(
		[]byte(json),
		v,
	)
	if err != nil {
		panic(err)
	}

	y := formats.NewYAML()
	yaml, err := y.Marshal(v)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(yaml))
}
```

``` console
$ go run ./example/json-to-yaml/json-to-yaml.go
name: Danny
roles:
- name: warrior
- name: worker

$ go run ./example/json-to-toml/json-to-toml.go
name = "Danny"

[[roles]]
name = "warrior"

[[roles]]
name = "worker"

```

## Limitations

There is a compatibility layer for:

- `JSON`, which helps to [mitigate](https://github.com/go-yaml/yaml/issues/139) non string keys in maps before they will be marshaled into `JSON`

### YAML

[go-yaml](https://github.com/go-yaml/yaml) handles struct keys without tags in [non-standard way](https://github.com/go-yaml/yaml/issues/148), they are lowercased.

There is no good workaround for this at the time of writing. Make sure you have tags for your struct field.

> Actualy I'd like to switch to more configurable yaml marshaler in the future, but at this time there is nothing better :(

### TOML

[toml](https://github.com/naoina/toml) can not marshal `interface{}` values at this time(panics, requires struct or map).

> Which is strange, probably some reflection misuse.

## License

MIT
