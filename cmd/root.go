package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tensorleap/leap-cli/cmd/auth"
	"github.com/tensorleap/leap-cli/cmd/cli"
	"github.com/tensorleap/leap-cli/cmd/mcp"
	"github.com/tensorleap/leap-cli/cmd/projects"
	"github.com/tensorleap/leap-cli/cmd/root_cmd"
	"github.com/tensorleap/leap-cli/cmd/run"
	"github.com/tensorleap/leap-cli/cmd/secrets"
	"github.com/tensorleap/leap-cli/cmd/server"
	"github.com/tensorleap/leap-cli/pkg/api"
	authPkg "github.com/tensorleap/leap-cli/pkg/auth"
	"github.com/tensorleap/leap-cli/pkg/config"
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
	SilenceUsage:  true,
	SilenceErrors: true,
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
	RootCommand.AddCommand(root_cmd.NewPushCmd())
	RootCommand.AddCommand(root_cmd.NewInitCmd())
	RootCommand.AddCommand(root_cmd.NewPullCmd())
	RootCommand.AddCommand(newVersionCmd())
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
	RootCommand.AddCommand(secrets.RootCommand)
	RootCommand.AddCommand(cli.RootCommand)
	RootCommand.AddCommand(projects.RootCommand)
	RootCommand.AddCommand(run.RootCommand)
	RootCommand.AddCommand(mcp.NewMcpCmd())

	initConfig()
}

// unknownSubcommand reports a subcommand a command group doesn't have. cobra does that only for the
// root: a group prints its help and exits 0, so `leap server check` looked like a check that passed
func unknownSubcommand(osArgs []string) error {
	cmd, args, err := RootCommand.Find(osArgs)
	if err != nil || cmd == RootCommand || !cmd.HasSubCommands() || cmd.Args != nil {
		return nil
	}
	if cmd.ParseFlags(args) != nil || cmd.Flags().NArg() == 0 {
		return nil
	}
	return fmt.Errorf("unknown command %q for %q\nRun '%s --help' for usage", cmd.Flags().Arg(0), cmd.CommandPath(), cmd.CommandPath())
}

func Execute() {
	err := unknownSubcommand(os.Args[1:])
	if err == nil {
		err = RootCommand.Execute()
	}
	if err != nil {
		if err != mcp.ErrShown {
			log.Error(api.HintVersionMismatch(err))
		}
		os.Exit(1)
	}
}

func initConfig() {
	err := config.InitConfig(cfgFile)
	cobra.CheckErr(err)
}
