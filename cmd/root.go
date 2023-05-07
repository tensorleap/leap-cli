package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tensorleap/cli-go/cmd/auth"
	"github.com/tensorleap/cli-go/cmd/local"
  k3d "github.com/k3d-io/k3d/v5/cmd"
)

var RootCmd = &cobra.Command{
	Use:   "leap",
	Short: "Tensorleap - Deepbug your models!",
	Long: `A debugger and analyzer for your DNNs.
Complete documentation is available at http://docs.tensoleap.ai`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

func init() {
	RootCmd.AddCommand(auth.AuthCommand)
	RootCmd.AddCommand(local.LocalCommand)
  RootCmd.AddCommand(k3d.NewCmdK3d())
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
