package auth

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/analytics"
	"github.com/tensorleap/leap-cli/pkg/api"
	"github.com/tensorleap/leap-cli/pkg/auth"
	"github.com/tensorleap/leap-cli/pkg/hint"
)

func NewLoginCmd() *cobra.Command {

	var name string
	var userName string
	var password string
	var apiKey string
	var nonInteractive bool

	cmd := &cobra.Command{
		Use:   "login [environment url]",
		Short: "Login using API key or username/password",
		Long: `Login to the system using either an API key or a username and password.

With --non-interactive it asks nothing, for scripts and AI assistants: it uses --api-key or
--username/--password when given, otherwise it prints a sign-in link to open in a browser
and waits until you have signed in there.`,
		Example: "  leap auth login http://localhost:4589\n" +
			"  leap auth login http://localhost:4589 --non-interactive",
		RunE: func(cmd *cobra.Command, args []string) error {
			var baseUrl string
			var err error
			var loginMethod string

			hasJustUrl := len(args) == 1
			if nonInteractive {
				if loginMethod, err = nonInteractiveMethod(hasJustUrl, apiKey, userName, password); err != nil {
					return err
				}
			}
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
					analytics.SendEvent(analytics.EventAuthLoginFailed, properties)
					return err
				}
			}
			apiUrl, err := api.NormalizeAPIUrl(baseUrl)
			if err != nil {
				// Track failed login attempt
				properties := map[string]interface{}{
					"api_url": baseUrl,
					"error":   err.Error(),
					"success": false,
					"method":  "url_validation",
				}
				analytics.SendEvent(analytics.EventAuthLoginFailed, properties)
				return err
			}

			hasUserNameOrPassword := userName != "" || password != ""
			hasApiKey := apiKey != ""
			askIfUseLogin := !hasUserNameOrPassword && !hasApiKey && !nonInteractive
			useLogin := hasUserNameOrPassword || loginMethod == "browser_oauth"
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
					analytics.SendEvent(analytics.EventAuthLoginFailed, properties)
					return err
				}
			}

			if useLogin {
				useBrowser := loginMethod == "browser_oauth"
				if !nonInteractive {
					useBrowser, err = auth.AskIfOpenBrowser()
					if err != nil {
						// Track failed login attempt
						properties := map[string]interface{}{
							"api_url": apiUrl,
							"error":   err.Error(),
							"success": false,
							"method":  "browser_selection",
						}
						analytics.SendEvent(analytics.EventAuthLoginFailed, properties)
						return err
					}
				}
				if useBrowser {
					loginMethod = "browser_oauth"
					apiKey, err = auth.LoginAndGetAuthTokenWithBrowser(cmd.Context(), apiUrl, !nonInteractive)
					if err != nil {
						// Track failed login attempt
						properties := map[string]interface{}{
							"api_url": apiUrl,
							"error":   err.Error(),
							"success": false,
							"method":  loginMethod,
						}
						analytics.SendEvent(analytics.EventAuthLoginFailed, properties)
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

						analytics.SendEvent(analytics.EventAuthLoginFailed, properties)
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

						analytics.SendEvent(analytics.EventAuthLoginFailed, properties)
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
					analytics.SendEvent(analytics.EventAuthLoginFailed, properties)
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
				analytics.SendEvent(analytics.EventAuthLoginFailed, properties)
				return err
			}

			if err := auth.Login(env); err != nil {
				return err
			}
			hint.Show(cmd.OutOrStdout(), "Next", hint.Integrate, hint.Pushed)
			return nil

		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "", "Name of the environment to login to, by default parse the api url")
	cmd.Flags().StringVarP(&apiKey, "api-key", "k", "", "API key to login with")
	cmd.Flags().StringVarP(&userName, "username", "u", "", "In case you want to login with username and password")
	cmd.Flags().StringVarP(&password, "password", "p", "", "In case you want to login with username and password")
	cmd.Flags().BoolVar(&nonInteractive, "non-interactive", false, "Ask nothing: use --api-key or --username/--password if given, otherwise print a browser sign-in link and wait for it")
	return cmd
}

// nonInteractiveMethod picks the login that asks nothing: the credentials given as flags, else the browser link
func nonInteractiveMethod(hasUrl bool, apiKey, userName, password string) (string, error) {
	switch {
	case !hasUrl:
		return "", fmt.Errorf("--non-interactive needs the server URL, e.g. leap auth login http://localhost:4589 --non-interactive")
	case apiKey != "":
		return "api_key_flag", nil
	case (userName == "") != (password == ""):
		return "", fmt.Errorf("--non-interactive needs both --username and --password, or neither to sign in with a browser link")
	case userName != "":
		return "username_password", nil
	}
	return "browser_oauth", nil
}

func init() {
	RootCommand.AddCommand(NewLoginCmd())
}
