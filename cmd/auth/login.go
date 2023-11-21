package auth

import (
	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/auth"
)

func NewLoginCmd() *cobra.Command {

	var name string

	cmd := &cobra.Command{
		Use:   "login [api key] [api url]",
		Args:  cobra.ExactArgs(2),
		Short: "Logging in with Tensorleap API key",
		Long:  `Logging in with Tensorleap API key`,
		RunE: func(cmd *cobra.Command, args []string) error {
			apiKey, baseUrl := args[0], args[1]

			if name == "" {
				name = auth.EnvNameFromUrl(baseUrl)
			}

			env := auth.NewEnv(name, baseUrl, apiKey)

			err := env.PrintWhoami(cmd.Context())
			if err != nil {
				return err
			}

			return auth.Login(env)

		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "", "Name of the environment to login to, by default parse the api url")
	return cmd
}

func init() {
	RootCommand.AddCommand(NewLoginCmd())
}
