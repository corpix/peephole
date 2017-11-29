package main

import (
	"io/ioutil"
	"os"

	"github.com/corpix/formats"
	"github.com/corpix/formats/compatibility"
	"github.com/urfave/cli"
)

func Action(ctx *cli.Context) error {
	var (
		v    = *new(interface{})
		buf  []byte
		from formats.Format
		to   formats.Format
		err  error
	)

	from, err = formats.New(ctx.String("from"))
	if err != nil {
		return err
	}

	to, err = formats.New(ctx.String("to"))
	if err != nil {
		return err
	}

	buf, err = ioutil.ReadAll(os.Stdin)
	if err != nil {
		return err
	}

	err = from.Unmarshal(buf, &v)
	if err != nil {
		return err
	}

	switch ctx.String("to") {
	case formats.JSON:
		v = compatibility.JSON(v)
	}

	buf, err = to.Marshal(v)
	if err != nil {
		return err
	}

	_, err = os.Stdout.Write(buf)
	if err != nil {
		return err
	}

	_, err = os.Stdout.Write([]byte{'\n'})

	return err
}

func main() {
	var (
		app = cli.NewApp()
	)

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "from",
			Value: "json",
			Usage: "From `what` format to encode",
		},
		cli.StringFlag{
			Name:  "to",
			Value: "json",
			Usage: "Decode into `specified` format",
		},
	}

	app.Action = Action

	app.RunAndExitOnError()
}
