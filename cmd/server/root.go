package server

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tensorleap/helm-charts/cmd/server"
	"github.com/tensorleap/helm-charts/pkg/local"
	serverpkg "github.com/tensorleap/helm-charts/pkg/server"
	"github.com/tensorleap/helm-charts/pkg/version"
)

var showInfo bool

var RootCommand = &cobra.Command{
	Use:   "server",
	Short: "Manage local server installation of Tensorleap",
	Long:  `Manage local server installation of Tensorleap`,
	Run: func(cmd *cobra.Command, args []string) {
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

func init() {
	serverpkg.SetInitDataDirFunc(initDataDir)
	RootCommand.PersistentFlags().BoolVarP(&showInfo, "info", "i", false, "Show server info")
	RootCommand.AddCommand(NewInstallCmd())
	RootCommand.AddCommand(NewUpgradeCmd())
	RootCommand.AddCommand(NewReinstallCmd())
	RootCommand.AddCommand(NewUninstallCmd())
	RootCommand.AddCommand(server.NewRunCmd())
	RootCommand.AddCommand(server.NewStopCmd())
	RootCommand.AddCommand(server.NewToolCmd())
}
