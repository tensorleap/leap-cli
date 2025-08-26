package auth

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/analytics"
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
			var baseUrl string
			var err error
			var loginMethod string

			hasJustUrl := len(args) == 1
			if hasJustUrl {
				baseUrl = args[0]
			} else {
				baseUrl, err = auth.AskForUrl("")
				if err != nil {
					// Track failed login attempt
					properties := map[string]interface{}{
						"error":   err.Error(),
						"success": false,
						"method":  "url_input",
					}
					_ = analytics.SendEvent(analytics.EventAuthLoginFailed, properties)
					return err
				}
			}
			apiUrl, err := auth.NormalizeAPIUrl(baseUrl)
			if err != nil {
				// Track failed login attempt
				properties := map[string]interface{}{
					"api_url": baseUrl,
					"error":   err.Error(),
					"success": false,
					"method":  "url_validation",
				}
				_ = analytics.SendEvent(analytics.EventAuthLoginFailed, properties)
				return err
			}

			hasUserNameOrPassword := userName != "" || password != ""
			hasApiKey := apiKey != ""
			askIfUseLogin := !hasUserNameOrPassword && !hasApiKey
			useLogin := hasUserNameOrPassword
			if askIfUseLogin {
				useLogin, err = auth.AskIfUseLogin()
				if err != nil {
					// Track failed login attempt
					properties := map[string]interface{}{
						"api_url": apiUrl,
						"error":   err.Error(),
						"success": false,
						"method":  "method_selection",
					}
					_ = analytics.SendEvent(analytics.EventAuthLoginFailed, properties)
					return err
				}
			}

			if useLogin {
				useBrowser, err := auth.AskIfOpenBrowser()
				if err != nil {
					// Track failed login attempt
					properties := map[string]interface{}{
						"api_url": apiUrl,
						"error":   err.Error(),
						"success": false,
						"method":  "browser_selection",
					}
					_ = analytics.SendEvent(analytics.EventAuthLoginFailed, properties)
					return err
				}
				if useBrowser {
					loginMethod = "browser_oauth"
					apiKey, err = auth.LoginAndGetAuthTokenWithBrowser(cmd.Context(), apiUrl)
					if err != nil {
						// Track failed login attempt
						properties := map[string]interface{}{
							"api_url": apiUrl,
							"error":   err.Error(),
							"success": false,
							"method":  loginMethod,
						}
						_ = analytics.SendEvent(analytics.EventAuthLoginFailed, properties)
						return err
					}
				} else {
					loginMethod = "username_password"
					userName, password, err = auth.AskForUserNameAndPassword(userName, password)
					if err != nil {
						// Track failed login attempt
						properties := map[string]interface{}{
							"api_url": apiUrl,
							"error":   err.Error(),
							"success": false,
							"method":  loginMethod,
						}

						// Add username if available
						if userName != "" {
							properties["inserted_username"] = userName
						}

						_ = analytics.SendEvent(analytics.EventAuthLoginFailed, properties)
						return err
					}
					apiKey, err = auth.LoginAndGetAuthToken(apiUrl, userName, password)
					if err != nil {
						// Track failed login attempt
						properties := map[string]interface{}{
							"api_url": apiUrl,
							"error":   err.Error(),
							"success": false,
							"method":  loginMethod,
						}

						// Add username if available
						if userName != "" {
							properties["inserted_username"] = userName
						}

						_ = analytics.SendEvent(analytics.EventAuthLoginFailed, properties)
						return err
					}
				}
			} else if !hasApiKey {
				loginMethod = "api_key_input"
				apiKey, err = auth.AskForApiKey()
				if err != nil {
					// Track failed login attempt
					properties := map[string]interface{}{
						"api_url": apiUrl,
						"error":   err.Error(),
						"success": false,
						"method":  loginMethod,
					}
					_ = analytics.SendEvent(analytics.EventAuthLoginFailed, properties)
					return err
				}
			} else {
				loginMethod = "api_key_flag"
			}
			apiKey = strings.TrimSpace(apiKey)

			if name == "" {
				name = auth.EnvNameFromUrl(apiUrl)
			}

			env := auth.NewEnv(name, apiUrl, apiKey)

			err = env.PrintWhoami(cmd.Context())
			if err != nil {
				// Track failed login attempt
				properties := map[string]interface{}{
					"api_url":  apiUrl,
					"env_name": name,
					"error":    err.Error(),
					"success":  false,
					"method":   loginMethod,
				}
				_ = analytics.SendEvent(analytics.EventAuthLoginFailed, properties)
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
