package server

import (
	"github.com/spf13/cobra"
	"github.com/tensorleap/helm-charts/cmd/server"
)

var RootCommand = &cobra.Command{
	Use:   "server",
	Short: "Manage local server installation of Tensorleap",
	Long:  `Manage local server installation of Tensorleap`,
}

func init() {
	RootCommand.AddCommand(NewInstallCmd())
	RootCommand.AddCommand(NewUpgradeCmd())
	RootCommand.AddCommand(NewReinstallCmd())
	RootCommand.AddCommand(server.NewRunCmd())
	RootCommand.AddCommand(server.NewStopCmd())
	RootCommand.AddCommand(server.NewUninstallCmd())
	RootCommand.AddCommand(server.NewToolCmd())
}
