package main

import (
	"log"

	"github.com/spf13/cobra/doc"
	"github.com/tensorleap/cli-go/cmd"
)

func main() {
	err := doc.GenMarkdownTree(cmd.RootCommand, "./docs")
	if err != nil {
		log.Fatal(err)
	}
}

