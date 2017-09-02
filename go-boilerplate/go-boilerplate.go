package main

import (
	"runtime"

	"github.com/corpix/go-boilerplate/cli"
)

func init() { runtime.GOMAXPROCS(runtime.NumCPU()) }
func main() { cli.Execute() }
