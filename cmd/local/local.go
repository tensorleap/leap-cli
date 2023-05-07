package local

import (
	"github.com/spf13/cobra"
)

var LocalCommand = &cobra.Command{
	Use:   "local",
  Short: "Manage local installation of Tensorleap",
	Long:  `Manage local installation of Tensorleap`,
	Run: func(cmd *cobra.Command, args []string) {
    cmd.Usage();
	},
}

func init() {
  LocalCommand.AddCommand(LocalInstallCommand)
}
