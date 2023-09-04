package local

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"runtime"

	"github.com/tensorleap/leap-cli/pkg/log"
)

const STANDALONE_DIR = "/var/lib/tensorleap/standalone"

func InitStandaloneDir() error {
	_, err := os.Stat(STANDALONE_DIR)
	if os.IsNotExist(err) {
		log.Printf("Creating directory: %s (you may be asked to enter the root user password)", STANDALONE_DIR)
		mkdirCmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("sudo mkdir -p %s", STANDALONE_DIR))
		if err := mkdirCmd.Run(); err != nil {
			return err
		}

		log.Println("Setting directory permissions")
		chmodCmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("sudo chmod -R 777 %s", STANDALONE_DIR))
		if err := chmodCmd.Run(); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	return initStandaloneSubDirs()
}

const (
	STORAGE_DIR_NAME  = "storage"
	REGISTRY_DIR_NAME = "registry"
	LOGS_DIR_NAME     = "logs"
)

func initStandaloneSubDirs() error {
	subDirs := []string{STORAGE_DIR_NAME, REGISTRY_DIR_NAME, LOGS_DIR_NAME}
	for _, dir := range subDirs {
		fullPath := path.Join(STANDALONE_DIR, dir)
		_, err := os.Stat(fullPath)
		if os.IsNotExist(err) {
			log.Printf("Creating directory: %s", fullPath)
			if err := os.MkdirAll(fullPath, 0777); err != nil {
				return err
			}
		} else if err != nil {
			return err
		}
	}
	return nil
}

func OpenLink(link string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", link)
	case "linux":
		cmd = exec.Command("xdg-open", link)
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", link)
	default:
		return fmt.Errorf("Unsupported platform!")
	}

	return cmd.Start()
}

func PurgeData() error {
	log.Infof("Purging data (you may be asked to enter the root user password)")
	for _, dir := range []string{STORAGE_DIR_NAME, REGISTRY_DIR_NAME} {
		path := path.Join(STANDALONE_DIR, dir)
		log.Infof("Removing directory: %s", path)
		err := os.RemoveAll(path)

		// if failed to remove directory, try to remove it with sudo
		if err != nil {
			rmCmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("sudo rm -rf %s", path))

			if err := rmCmd.Run(); err != nil {
				log.SendCloudReport("error", "Failed purge data", "Failed", &map[string]interface{}{"error": err.Error()})
				return err
			}
		}

	}
	return nil
}
