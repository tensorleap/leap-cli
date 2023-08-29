package server

import (
	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/k3d"
	"github.com/tensorleap/leap-cli/pkg/local"
	"github.com/tensorleap/leap-cli/pkg/log"
)

func NewRunCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "run",
		Aliases: []string{"up", "start"},
		Short:   "Run Tensorleap server",
		Long:    `Run Tensorleap server`,
		RunE: func(cmd *cobra.Command, args []string) error {
			log.SetCommandName("run")

			close, err := local.SetupInfra("run")
			if err != nil {
				return err
			}
			defer close()

			err = k3d.RunCluster(cmd.Context())
			if err != nil {
				return err
			}
			return nil
		},
	}
	return cmd
}

func init() {
	RootCommand.AddCommand(NewRunCmd())
}
