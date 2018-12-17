package alchemy

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/urfave/cli"
)

/**
 * Interface Declarations
 * ------------------------------------------------------------------------------------------------
 */

// CliApp represents an instance of a running Alchemy session. Mostly just a wrapper
// around the urfave/cli package and a few configuration values.
type CliApp struct {
	Io *cli.App
}

// Parameters represents the parsed, validated and evaluated CLI argument set
type Parameters struct {
	Indexes      []string
	ConfigValues AlchemyRC
}

// {
// 	"appId": "algolia app ID here",
// 	"searchKey": "algolia search key here",
// 	"secretKey": "algolia secret key here",
// 	"fixtures": "./fixtures.json",
// 	"tests": "./index_name.tests.json"
// }
// AlchemyRC represents our .alchemyrc configuration file
type AlchemyRC struct {
	AppID        string `json:"appId"`
	SearchKey    string `json:"searchKey"`
	SecretKey    string `json:"secretKey"`
	FixturesPath string `json:"fixtures"`
	TestsPath    string `json:"tests"`
}

/**
 * Local Setup
 * ------------------------------------------------------------------------------------------------
 */

var flags = []cli.Flag{
	cli.StringFlag{
		Name:  "config, c",
		Value: "./.alchemyrc",
		Usage: "Load configuration from `FILE`",
	},
}

/**
 * Public Methods
 * ------------------------------------------------------------------------------------------------
 */

// New creates a new instance of CliApp
func New(version string) *CliApp {
	io := cli.NewApp()
	io.Version = version
	io.Usage = "An acceptance testing tool for Algolia indexes"
	io.Flags = flags

	if len(os.Args) > 1 {
		io.Action = validate(Run)
	} else {
		io.Action = cli.ShowAppHelp
	}

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

/**
 * Local Methods
 * ------------------------------------------------------------------------------------------------
 */

// validate the input / CLI parameters
func validate(then func(params Parameters, c *cli.Context) error) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		// index name
		var indexes []string
		if c.NArg() == 1 {
			indexes = append(indexes, c.Args()[0])
		} else if c.NArg() > 1 {
			indexes = append(indexes, c.Args()...)
		}

		if len(indexes) > 1 {
			return errors.New("ERROR: Alchemy currently only supports running tests against a single index")
		}

		// config file
		configPath := path.Clean(c.String("config"))

		var configValues AlchemyRC

		_, err := os.Stat(configPath)
		if err != nil {
			return fmt.Errorf("Couldn't access config file '%s'", configPath)
		}

		configFile, err := os.Open(configPath)
		defer configFile.Close()
		if err != nil {
			return err
		}
		parser := json.NewDecoder(configFile)
		parser.Decode(&configValues)

		// Now we've passed validation, let's build up a Parameters object and call
		// our app run callback with the parameters (and the context, just in case)
		params := Parameters{
			Indexes:      indexes,
			ConfigValues: configValues,
		}

		return then(params, c)
	}
}
