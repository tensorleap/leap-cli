package cli

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/log"
)

const (
	INSTALL_SCRIPT_URL = "https://raw.githubusercontent.com/tensorleap/leap-cli/master/install.sh"
)

func NewUpgradeCliCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upgrade",
		Short: "Upgrades the tensorleap cli installation on the local machine",
		Long:  `Upgrades the tensorleap cli installation on the local machine`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := UpgradeCLI()
			if err != nil {
				fmt.Println("Error installing Leap CLI:", err)
			}

			return nil
		},
	}

	return cmd
}

func init() {
	RootCommand.AddCommand(NewUpgradeCliCmd())
}

func UpgradeCLI() error {
	// Download the script
	resp, err := http.Get(INSTALL_SCRIPT_URL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Read the contents of the script
	script, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Run the script
	cmd := exec.Command("bash", "-c", string(script))
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}
	err = cmd.Start()
	if err != nil {
		return err
	}

	// Log any output generated by the script
	go printOutput(stdout, log.Println)
	go printOutput(stderr, log.Error)

	err = cmd.Wait()
	if err != nil {
		return err
	}

	return nil
}

func printOutput(pipeStd io.ReadCloser, logger func(...interface{})) {
	scanner := bufio.NewScanner(pipeStd)
	for scanner.Scan() {
		logger(scanner.Text())
	}
}
