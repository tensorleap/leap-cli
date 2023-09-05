package cli

import (
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/cli"
)

const (
	INSTALL_SCRIPT_URL = "https://raw.githubusercontent.com/tensorleap/leap-cli/master/install.sh"
)

func NewUpgradeCliCmd() *cobra.Command {
	var installScript bool
	cmd := &cobra.Command{
		Use:   "upgrade",
		Short: "Helps to upgrade the leap cli",
		Long:  `Helps to upgrade the leap cli`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if !installScript {
				fmt.Println("Run the following command to upgrade:")
				fmt.Println(cli.UpgradeCmd)
				return nil
			}
			err := printScript()
			if err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().BoolVarP(&installScript, "script", "s", false, "Print the install script to stdout")

	return cmd
}

func init() {
	RootCommand.AddCommand(NewUpgradeCliCmd())
}

func printScript() error {
	resp, err := http.Get(INSTALL_SCRIPT_URL)
	if err != nil {
		return fmt.Errorf("failed to fetch install script: %v", err)
	}
	defer resp.Body.Close()

	script, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read install script: %v", err)
	}
	fmt.Print(string(script))
	return nil
}
