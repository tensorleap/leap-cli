package local

import (
	"github.com/spf13/cobra"
)

var LocalCommand = &cobra.Command{
	Use:   "local",
  Short: "Manage local installation of Tensorleap",
	Long:  `Manage local installation of Tensorleap`,
}

func init() {
  LocalCommand.AddCommand(LocalInstallCommand)
}
