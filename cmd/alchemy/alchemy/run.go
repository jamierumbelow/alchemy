package alchemy

import (
	"errors"
	"fmt"
	"strings"

	"github.com/urfave/cli"
)

// Run the main Alchemy tests
func Run(c *cli.Context) error {
	var indexes []string
	if c.NArg() == 1 {
		indexes = append(indexes, c.Args()[0])
	} else if c.NArg() > 1 {
		indexes = append(indexes, c.Args()...)
	}

	if len(indexes) > 1 {
		return errors.New("ERROR: Alchemy currently only supports running tests against a single index")
	}

	fmt.Printf("Pushing to indexes: %s", strings.Join(indexes, ","))
	return nil
}
