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
