package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tensorleap/leap-cli/cmd/auth"
	"github.com/tensorleap/leap-cli/cmd/cli"
	"github.com/tensorleap/leap-cli/cmd/code"
	"github.com/tensorleap/leap-cli/cmd/hub"
	"github.com/tensorleap/leap-cli/cmd/models"
	"github.com/tensorleap/leap-cli/cmd/projects"
	"github.com/tensorleap/leap-cli/cmd/run"
	"github.com/tensorleap/leap-cli/cmd/secrets"
	"github.com/tensorleap/leap-cli/cmd/server"
	authPkg "github.com/tensorleap/leap-cli/pkg/auth"
	"github.com/tensorleap/leap-cli/pkg/config"
	hubPkg "github.com/tensorleap/leap-cli/pkg/hub"
	"github.com/tensorleap/leap-cli/pkg/log"
	"github.com/tensorleap/leap-cli/pkg/version"
)

var cfgFile string
var apiKey string
var apiUrl string
var showVersionInfo bool

var RootCommand = &cobra.Command{
	Use:   "leap",
	Short: "Leap - Deepbug your models!",
	Long: `A debugger and analyzer for your DNNs.
Complete documentation is available at https://docs.tensorleap.ai`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		env := authPkg.GetCurrentEnv()
		authCtx := env.AuthContext(
			context.Background(),
		)
		cmd.SetContext(authCtx)
	},
	Run: func(cmd *cobra.Command, args []string) {
		if !showVersionInfo {
			_ = cmd.Help()
			return
		}

		fmt.Println(version.CliVersion)
	},
}

func init() {

	RootCommand.Flags().BoolVar(&showVersionInfo, "version", false, "Show version information")
	RootCommand.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/tensorleap/config.yaml)")
	RootCommand.PersistentFlags().StringVar(&apiKey, "apiKey", "", "Tensorleap Api key")
	RootCommand.PersistentFlags().StringVar(&apiUrl, "apiUrl", "", "Tensorleap api url")
	if err := viper.BindPFlag(authPkg.API_KEY_CONFIG_PATH, RootCommand.PersistentFlags().Lookup("apiKey")); err != nil {
		log.Error(err)
	}
	if err := viper.BindPFlag(authPkg.API_URL_CONFIG_PATH, RootCommand.PersistentFlags().Lookup("apiUrl")); err != nil {
		log.Error(err)
	}
	RootCommand.AddCommand(auth.RootCommand)
	RootCommand.AddCommand(server.RootCommand)
	RootCommand.AddCommand(code.RootCommand)
	RootCommand.AddCommand(secrets.RootCommand)
	RootCommand.AddCommand(models.RootCommand)
	RootCommand.AddCommand(cli.RootCommand)
	RootCommand.AddCommand(projects.RootCommand)
	RootCommand.AddCommand(run.RootCommand)
	if hubPkg.IsHubEnabled() {
		RootCommand.AddCommand(hub.RootCommand)
	}

	initConfig()
}

func Execute() {
	if err := RootCommand.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}

func initConfig() {
	err := config.InitConfig(cfgFile)
	cobra.CheckErr(err)
}
