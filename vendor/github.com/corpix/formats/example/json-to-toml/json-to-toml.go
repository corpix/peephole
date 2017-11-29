package main

import (
	"fmt"

	"github.com/corpix/formats"
)

var (
	json = `
        {
            "name": "Danny",
            "roles": [{"name": "warrior"}, {"name": "worker"}]
        }
    `
)

func main() {
	// XXX: map[string]interface{} because
	// toml package says:
	// panic: toml: cannot marshal interface {} as table, want struct or map type
	v := new(interface{})

	j := formats.NewJSON()
	err := j.Unmarshal(
		[]byte(json),
		v,
	)
	if err != nil {
		panic(err)
	}

	t := formats.NewTOML()
	toml, err := t.Marshal(v)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(toml))
}
