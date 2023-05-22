package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tensorleap/cli-go/cmd/auth"
	"github.com/tensorleap/cli-go/cmd/datasets"
	"github.com/tensorleap/cli-go/cmd/local"
	"github.com/tensorleap/cli-go/cmd/models"
	. "github.com/tensorleap/cli-go/pkg/api"
	"github.com/tensorleap/cli-go/pkg/config"
)

var RootCommand = &cobra.Command{
	Use:   "tensorleap",
	Short: "Tensorleap - Deepbug your models!",
	Long: `A debugger and analyzer for your DNNs.
Complete documentation is available at http://docs.tensoleap.ai`,
  PersistentPreRun: func(cmd *cobra.Command, args []string) {
    cmd.SetContext(CreateAuthenticatedContext(
      cmd.Context(),
      config.GetApiUrl(),
      config.GetApiKey()))
  },
}

var cfgFile string
var apiKey string
var apiUrl string

func init() {
  cobra.OnInitialize(initConfig)

	RootCommand.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/tensorleap/config.yaml)")
	RootCommand.PersistentFlags().StringVar(&apiKey, "apiKey", "", "Tensorleap Api key")
	RootCommand.PersistentFlags().StringVar(&apiUrl, "apiUrl", "", "Tensorleap api url")
  viper.BindPFlag(config.API_KEY_CONFIG_PATH, RootCommand.PersistentFlags().Lookup("apiKey"))
  viper.BindPFlag(config.API_URL_CONFIG_PATH, RootCommand.PersistentFlags().Lookup("apiUrl"))

	RootCommand.AddCommand(auth.RootCommand)
	RootCommand.AddCommand(local.RootCommand)
  RootCommand.AddCommand(datasets.RootCommand)
  RootCommand.AddCommand(models.RootCommand)
}

func Execute() {
	if err := RootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initConfig() {
  config.InitConfig(cfgFile)
}

