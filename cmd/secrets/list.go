package secrets

import (
	"github.com/tensorleap/leap-cli/pkg/entity"
	"github.com/tensorleap/leap-cli/pkg/secret"

	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/auth"
)

func NewListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List secrets",
		Long:  `List secrets`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := auth.CheckLoggedIn(); err != nil {
				return err
			}
			list, err := secret.GetSecretManagerList(cmd.Context())
			if err != nil {
				return err
			}
			entity.PrintList(list, secret.SecretEntityDesc)

			return nil
		},
	}

	return cmd
}

func init() {
	RootCommand.AddCommand(NewListCmd())
}
