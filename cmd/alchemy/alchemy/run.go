package alchemy

import (
	"fmt"
	"strings"

	"github.com/urfave/cli"
)

// Run the main Alchemy tests
func Run(params Parameters, c *cli.Context) error {
	fmt.Print(params)
	fmt.Printf("Pushing to indexes: %s", strings.Join(params.Indexes, ","))

	return nil
}
