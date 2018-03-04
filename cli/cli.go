package cli

import (
	"fmt"
	builtinLogger "log"
	"os"

	"github.com/corpix/loggers"
	"github.com/davecgh/go-spew/spew"
	"github.com/urfave/cli"

	"github.com/corpix/go-boilerplate/config"
	appLogger "github.com/corpix/go-boilerplate/logger"
)

var (
	version = "development"

	// log is a application-wide logger.
	log loggers.Logger

	// Config is a container that represents the current application configuration.
	Config config.Config
)

// Prerun configures application before running and executing from urfave/cli.
func Prerun(c *cli.Context) error {
	var err error

	err = initConfig(c)
	if err != nil {
		return err
	}

	err = initLogger(c)
	if err != nil {
		return err
	}

	if c.Bool("debug") {
		fmt.Println("Dumping configuration structure")
		spew.Dump(Config)
	}

	return nil
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	app := cli.NewApp()
	app.Name = "go-boilerplate"
	app.Usage = "go-boilerplate-description"
	app.Action = RootAction
	app.Flags = RootFlags
	app.Commands = RootCommands
	app.Before = Prerun
	app.Version = version

	err := app.Run(os.Args)
	if err != nil {
		builtinLogger.Fatal(err)
	}
}

// initConfig reads in config file.
func initConfig(ctx *cli.Context) error {
	var (
		path = os.ExpandEnv(ctx.String("config"))
		err  error
	)

	Config, err = config.FromFile(path)
	if err != nil {
		return err
	}

	return nil
}

// initLogger inits logger component with
// parameters from config.
func initLogger(c *cli.Context) error {
	var (
		err error
	)

	if c.Bool("debug") {
		Config.Logger.Level = "debug"
	}

	log, err = appLogger.FromConfig(Config.Logger)
	if err != nil {
		return err
	}

	return nil
}
