package cli

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"

	"github.com/tensorleap/cli-go/pkg/log"

	"github.com/spf13/cobra"
)

const (
	REPO_URL = "https://api.github.com/repos/tensorleap/cli-go/releases/latest"
	APP_NAME = "tensorleap"
)

type Release struct {
	TagName string `json:"tag_name"`
}

func NewUpgradeCliCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upgrade",
		Short: "Upgrades the tensorleap cli installation on the local machine",
		Long:  `Upgrades the tensorleap cli installation on the local machine`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := UpgradeCLI()
			if err != nil {
				fmt.Println("Error installing TensorLeap CLI:", err)
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
	cmd := exec.Command("bash", "-c", "wget -q -O - https://raw.githubusercontent.com/tensorleap/cli-go/master/install.sh | bash")

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

	go printOutput(stdout, log.Println)
	go printOutput(stderr, log.Error)

	err = cmd.Wait()
	if err != nil {
		return fmt.Errorf("command failed: %v", err)
	}

	return nil
}

func printOutput(pipeStd io.ReadCloser, logger func(...interface{})) {
	scanner := bufio.NewScanner(pipeStd)
	for scanner.Scan() {
		logger(scanner.Text())
	}
}
