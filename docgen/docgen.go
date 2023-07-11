package main

import (
	"github.com/spf13/cobra/doc"
	"github.com/tensorleap/cli-go/cmd"
	"github.com/tensorleap/cli-go/pkg/log"
)

func main() {
	err := doc.GenMarkdownTree(cmd.RootCommand, "./docs")
	if err != nil {
		log.Fatalln(err)
	}
}
