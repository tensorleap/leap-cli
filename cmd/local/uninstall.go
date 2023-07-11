package local

import (
	"github.com/spf13/cobra"
	"github.com/tensorleap/cli-go/pkg/k3d"
	"github.com/tensorleap/cli-go/pkg/local"
)


func NewUninstallCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "uninstall",
		Short: "Remove local Tensorleap installation",
		Long:  `Remove local Tensorleap installation`,
		RunE: func(cmd *cobra.Command, args []string) error {
			local.SetupBackgroundLogger("uninstall")
			ctx := cmd.Context()
			err := k3d.UninstallCluster(ctx)
			return err

			// todo: implement --cache and --data
		},
	}
	return cmd
}

func init() {
	RootCommand.AddCommand(NewUninstallCmd())
}
