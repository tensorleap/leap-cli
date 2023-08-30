package secrets

import (
	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/secret"
)

func NewCreateCmd() *cobra.Command {
	var key string

	cmd := &cobra.Command{
		Use:   "create [name] [secretKeyPath]",
		Short: "Create a new secret",
		Long:  `Create a new secret`,
		RunE: func(cmd *cobra.Command, args []string) error {
			var name string
			var keyPath string
			if len(args) > 0 {
				name = args[0]
			}
			if len(args) > 1 {
				keyPath = args[1]
			}

			secrets, err := secret.GetSecretManagerList(cmd.Context())
			if err != nil {
				return err
			}

			err = secret.AskForSecret(cmd.Context(), secrets, &name, &key, keyPath)
			if err != nil {
				return err

			}

			_, err = secret.CreateSecretManager(cmd.Context(), name, key)
			if err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().StringVarP(&key, "secret-key-content", "k", "", "Secret key content")

	return cmd
}

func init() {
	RootCommand.AddCommand(NewCreateCmd())
}
