package main

import (
	"github.com/spf13/cobra/doc"
	"github.com/tensorleap/leap-cli/cmd"
	"github.com/tensorleap/leap-cli/pkg/log"
)

func main() {
	err := doc.GenMarkdownTree(cmd.RootCommand, "./docs")
	if err != nil {
		log.Fatalln(err)
	}
}
