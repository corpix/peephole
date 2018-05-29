package cli

import (
	"fmt"
	"time"

	"github.com/urfave/cli"

	"github.com/corpix/peephole/proxy"
	"github.com/corpix/peephole/proxy/metrics"
	"github.com/corpix/peephole/socks"
)

var (
	// RootCommands is a list of subcommands for the application.
	RootCommands = []cli.Command{}

	// RootFlags is a list of flags for the application.
	RootFlags = []cli.Flag{
		cli.StringFlag{
			Name:   "config, c",
			Usage:  "application configuration file",
			EnvVar: "CONFIG",
			Value:  "config.json",
		},
		cli.BoolFlag{
			Name:   "debug",
			EnvVar: "DEBUG",
			Usage:  "add this flag to enable debug mode",
		},
		cli.BoolFlag{
			Name:  "profile",
			Usage: "write profile information for debugging(cpu.prof, heap.prof)",
		},
		cli.BoolFlag{
			Name:  "trace",
			Usage: "write trace information for debugging(trace.prof)",
		},
	}
)

// RootAction is executing when program called without any subcommand.
func RootAction(c *cli.Context) error {
	m, err := metrics.New(Config.Proxy.Metrics, log)
	if err != nil {
		log.Fatal(err)
	}

	p, err := proxy.NewParams(Config.Proxy, log)
	if err != nil {
		log.Fatal(err)
	}

	server := socks.New(p)

	for {
		err = server.ListenAndServe("tcp", Config.Listen)
		if err != nil {
			log.Error(err)

			m.IncrCounter(
				[]string{"errors", fmt.Sprintf("%T", err)},
				1,
			)
		}

		time.Sleep(5 * time.Second)
	}
}
