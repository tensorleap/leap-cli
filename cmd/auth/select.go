package auth

import (
	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/auth"
)

func NewSelectCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "select [environment name]",
		Short: "Select an environment to use",
		Long:  `Select an environment to use`,
		RunE: func(cmd *cobra.Command, args []string) error {
			var envNameArg string
			if len(args) > 0 {
				envNameArg = args[0]
			}

			return auth.SelectAndSetEnv(envNameArg)
		},
	}
	return cmd
}

func init() {
	RootCommand.AddCommand(NewSelectCmd())
}
