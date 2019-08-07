package main

import (
	"log"

	"github.com/pulkitsharma07/go-cli-boilerplate/cmd"
	"github.com/spf13/cobra/doc"
)

func main() {
	err := doc.GenMarkdownTree(cmd.Command(), "./")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Docs generated !")
}
