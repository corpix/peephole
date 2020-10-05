package main

import (
	"runtime"

	"github.com/corpix/peephole/cli"
)

func init() { runtime.GOMAXPROCS(runtime.NumCPU()) }
func main() { cli.Run() }
