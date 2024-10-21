package server

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tensorleap/helm-charts/pkg/local"
	"github.com/tensorleap/helm-charts/pkg/version"
)

var showInfo bool
var isNotSupported bool

var RootCommand = &cobra.Command{
	Use:   "server",
	Short: "Manage local server installation of Tensorleap",
	Long:  `Manage local server installation of Tensorleap`,
	Run: func(cmd *cobra.Command, args []string) {
		if isNotSupported {
			fmt.Println("This command is not supported on this platform")
			return
		}
		if showInfo {
			dataDir := getConfigureDataDir()
			if dataDir == "" {
				dataDir = local.GetServerDataDir()
			}
			fmt.Println("Installer Version: ", version.Version)
			fmt.Println("Data dir:          ", dataDir)

		} else {
			_ = cmd.Help()
		}
	},
}
