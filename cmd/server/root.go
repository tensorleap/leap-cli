package server

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tensorleap/helm-charts/cmd/server"
	"github.com/tensorleap/helm-charts/pkg/version"
)

var showVersion bool

var RootCommand = &cobra.Command{
	Use:   "server",
	Short: "Manage local server installation of Tensorleap",
	Long:  `Manage local server installation of Tensorleap`,
	Run: func(cmd *cobra.Command, args []string) {
		if showVersion {
			fmt.Println(version.Version)
		} else {
			_ = cmd.Help()
		}
	},
}

func init() {
	RootCommand.PersistentFlags().BoolVarP(&showVersion, "version", "v", false, "Show server installer version")
	RootCommand.AddCommand(NewInstallCmd())
	RootCommand.AddCommand(NewUpgradeCmd())
	RootCommand.AddCommand(NewReinstallCmd())
	RootCommand.AddCommand(server.NewRunCmd())
	RootCommand.AddCommand(server.NewStopCmd())
	RootCommand.AddCommand(server.NewUninstallCmd())
	RootCommand.AddCommand(server.NewToolCmd())
	RootCommand.AddCommand(NewHostnameCmd())
}
