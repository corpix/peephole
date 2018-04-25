package cli

import (
	"time"

	"github.com/urfave/cli"

	"github.com/corpix/peephole/proxy"
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
	cfg, err := proxy.NewConfig(Config, log)
	if err != nil {
		log.Fatal(err)
	}

	server, err := socks.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	for {
		err = server.ListenAndServe("tcp", Config.Addr)
		if err != nil {
			log.Error(err)
		}

		time.Sleep(5 * time.Second)
	}
}
