package local

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"

	"github.com/spf13/cobra"

	"github.com/tensorleap/cli-go/pkg/helm"
	"github.com/tensorleap/cli-go/pkg/k3d"
)

const VAR_DIR = "/var/lib/tensorleap/standalone"

var port uint
var registryPort uint

func init() {
	cmd := &cobra.Command{
		Use:   "install",
		Short: "Installs tensorleap on the local machine, running in a docker container",
		Long:  `Installs tensorleap on the local machine, running in a docker container`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := initVarDir(); err != nil {
				return err
			}

			ctx := cmd.Context()

			regVolumeDir := path.Join(VAR_DIR, "registry")
			if err := k3d.CreateRegistry(
				ctx,
				registryPort,
				[]string{fmt.Sprintf("%v:%v", regVolumeDir, "/var/lib/registry")},
			); err != nil {
				return err
			}

			if err := k3d.CreateCluster(
				ctx,
				port,
				[]string{fmt.Sprintf("%v:%v", VAR_DIR, VAR_DIR)},
			); err != nil {
				return err
			}

			if err := helm.InstallLatestTensorleapChartVersion(
				ctx,
				"k3d-tensorleap",
				"tensorleap",
			); err != nil {
				return err
			}

			return nil

		},
	}

	cmd.Flags().UintVarP(&port, "port", "p", 4589, "Port to be used for tensorleap installation")
	cmd.Flags().UintVar(&registryPort, "registry-port", 5699, "Port to be used for docker registyr")

	RootCommand.AddCommand(cmd)
}

func initVarDir() error {
	_, err := os.Stat(VAR_DIR)
	if os.IsNotExist(err) {
		log.Printf("Creating directory: %s (you may be asked to enter the root user password)", VAR_DIR)
		mkdirCmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("sudo mkdir -p %s", VAR_DIR))
		if err := mkdirCmd.Run(); err != nil {
			return err
		}

		log.Print("Setting directory permissions")
		chmodCmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("sudo chmod -R 777 %s", VAR_DIR))
		if err := chmodCmd.Run(); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	for _, dir := range []string{"manifests", "storage", "registry"} {
		fullPath := path.Join(VAR_DIR, dir)
		_, err := os.Stat(fullPath)
		if os.IsNotExist(err) {
			log.Printf("Creating directory: %s", fullPath)
			if err := os.MkdirAll(fullPath, 777); err != nil {
				return err
			}
		} else if err != nil {
			return err
		}
	}

	return nil
}
