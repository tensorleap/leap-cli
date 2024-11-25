//go:build !windows
// +build !windows

package server

import (
	"github.com/tensorleap/helm-charts/cmd/server"
	serverpkg "github.com/tensorleap/helm-charts/pkg/server"
)

func init() {
	serverpkg.SetInitDataDirFunc(initDataDir)
	RootCommand.PersistentFlags().BoolVarP(&showInfo, "info", "i", false, "Show server info")
	RootCommand.AddCommand(NewInstallCmd())
	RootCommand.AddCommand(NewUpgradeCmd())
	RootCommand.AddCommand(NewReinstallCmd())
	RootCommand.AddCommand(NewUninstallCmd())
	RootCommand.AddCommand(NewInfoCmd())
	RootCommand.AddCommand(server.NewRunCmd())
	RootCommand.AddCommand(server.NewStopCmd())
	RootCommand.AddCommand(server.NewToolCmd())
}
