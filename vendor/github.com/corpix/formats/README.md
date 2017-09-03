formats
----------

[![Build Status](https://travis-ci.org/corpix/formats.svg?branch=master)](https://travis-ci.org/corpix/formats)

This package provides a consistent API to transform same data(represented by some `struct`) between arbitrary formats.

Supported formats:

- `JSON`
- `YAML`

## Example

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
$ go run ./example/json-to-yaml.go
name: Danny
roles:
- warrior
- worker
```

## Compatibility

There is a compatibility layer for:

- `JSON`, which helps to [mitigate](https://github.com/go-yaml/yaml/issues/139) non string keys in maps before they will be marshaled into `JSON`

### YAML

[go-yaml](https://github.com/go-yaml/yaml) handles struct keys without tags in [non-standard way](https://github.com/go-yaml/yaml/issues/148), they are lowercased.

There is no good workaround for this at the time of writing. Make sure you have tags for your struct field.

> Actualy I'd like to switch to more configurable yaml marshaler in the future, but at this time there is nothing better :(

## License

MIT
