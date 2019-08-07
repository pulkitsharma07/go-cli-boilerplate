package main

import (
	"github.com/pulkitsharma07/go-cli-boilerplate/cmds"
	"os"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
