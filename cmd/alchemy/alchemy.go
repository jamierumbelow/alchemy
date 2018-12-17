package main

import (
	"github.com/jamierumbelow/alchemy/cmd/alchemy/alchemy"
)

const version = "0.0.1"

func main() {
	app := alchemy.New(version)
	app.Run()
}
