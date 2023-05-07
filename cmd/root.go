package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tensorleap/cli-go/cmd/auth"
	"github.com/tensorleap/cli-go/cmd/local"
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
	RootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "Author name for copyright attribution")
	RootCmd.AddCommand(auth.AuthCommand)
	RootCmd.AddCommand(local.LocalCommand)
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
