package alchemy

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

// CliApp represents an instance of a running Alchemy session. Mostly just a wrapper
// around the urfave/cli package and a few configuration values.
type CliApp struct {
	Io *cli.App
}

var flags = []cli.Flag{
	cli.StringFlag{
		Name:  "config, c",
		Value: "./.alchemyrc",
		Usage: "Load configuration from `FILE`",
	},
}

// New creates a new instance of CliApp
func New(version string) *CliApp {
	io := cli.NewApp()
	io.Version = version
	io.Usage = "An acceptance testing tool for Algolia indexes"
	io.Flags = flags
	io.Action = Run

	cliApp := CliApp{Io: io}

	return &cliApp
}

// Run runs the application
func (app CliApp) Run() {
	err := app.Io.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
