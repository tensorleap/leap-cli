package auth

import (
	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/auth"
)

func NewLoginCmd() *cobra.Command {

	var name string
	var userName string
	var password string
	var apiKey string

	cmd := &cobra.Command{
		Use:   "login [environment url]",
		Short: "Login using API key or username/password",
		Long:  `Login to the system using either an API key or a username and password`,
		RunE: func(cmd *cobra.Command, args []string) error {
			var apiKey, baseUrl string
			var err error

			hasJustUrl := len(args) == 1
			if hasJustUrl {
				baseUrl = args[0]
			} else {
				baseUrl, err = auth.AskForUrl("")
				if err != nil {
					return err
				}
			}
			apiUrl := auth.NormalizeAPIUrl(baseUrl)

			hasUserNameOrPassword := userName != "" || password != ""
			hasApiKey := apiKey != ""
			askIfUseUserAndPassword := !hasUserNameOrPassword && !hasApiKey
			useLogin := hasUserNameOrPassword
			if askIfUseUserAndPassword {
				useLogin, err = auth.AskIfUseUserAndPassword()
				if err != nil {
					return err
				}
			}

			if useLogin {
				userName, password, err = auth.AskForUserNameAndPassword(userName, password)
				if err != nil {
					return err
				}
				apiKey, err = auth.LoginAndGetAuthToken(apiUrl, userName, password)
				if err != nil {
					return err
				}
			} else if !hasApiKey {
				apiKey, err = auth.AskForApiKey()
				if err != nil {
					return err
				}
			}

			if name == "" {
				name = auth.EnvNameFromUrl(apiUrl)
			}

			env := auth.NewEnv(name, apiUrl, apiKey)

			err = env.PrintWhoami(cmd.Context())
			if err != nil {
				return err
			}

			return auth.Login(env)

		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "", "Name of the environment to login to, by default parse the api url")
	cmd.Flags().StringVarP(&apiKey, "api-key", "k", "", "API key to login with")
	cmd.Flags().StringVarP(&userName, "username", "u", "", "In case you want to login with username and password")
	cmd.Flags().StringVarP(&password, "password", "p", "", "In case you want to login with username and password")
	return cmd
}

func init() {
	RootCommand.AddCommand(NewLoginCmd())
}
