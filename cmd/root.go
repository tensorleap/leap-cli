package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tensorleap/cli-go/cmd/auth"
	"github.com/tensorleap/cli-go/cmd/datasets"
	"github.com/tensorleap/cli-go/cmd/local"
	"github.com/tensorleap/cli-go/cmd/models"
)

var RootCommand = &cobra.Command{
	Use:   "leap",
	Short: "Tensorleap - Deepbug your models!",
	Long: `A debugger and analyzer for your DNNs.
Complete documentation is available at http://docs.tensoleap.ai`,
}

func init() {
	RootCommand.AddCommand(auth.RootCommand)
	RootCommand.AddCommand(local.RootCommand)
  RootCommand.AddCommand(datasets.RootCommand)
  RootCommand.AddCommand(models.RootCommand)
}

func Execute() {
	if err := RootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
