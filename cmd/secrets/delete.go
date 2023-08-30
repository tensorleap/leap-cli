package secrets

import (
	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/entity"
	"github.com/tensorleap/leap-cli/pkg/secret"
)

func NewDeleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete [name]",
		Short: "Delete a secret",
		Long:  `Delete a secret`,
		RunE: func(cmd *cobra.Command, args []string) error {
			list, err := secret.GetSecretManagerList(cmd.Context())
			if err != nil {
				return err
			}

			var selectedSecret *secret.SecretEntity
			if len(args) > 0 {
				name := args[0]
				selectedSecret, err = entity.GetEntityByDisplayName(name, list, secret.SecretEntityDesc)
				if err != nil {
					return err
				}
			} else {
				selectedSecret, err = entity.SelectEntity(list, secret.SecretEntityDesc)
				if err != nil {
					return err
				}
			}

			err = secret.DeleteSecretManager(cmd.Context(), selectedSecret)
			return err
		},
	}

	return cmd
}

func init() {
	RootCommand.AddCommand(NewDeleteCmd())
}
