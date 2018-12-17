package alchemy

import (
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

	fmt.Printf("Pushing to indexes: %s", strings.Join(indexes, ","))
	return nil
}
