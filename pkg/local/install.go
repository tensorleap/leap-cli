package local

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

func InitVarDir() error {
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

	for _, dir := range []string{"storage", "registry"} {
		fullPath := path.Join(VAR_DIR, dir)
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

func InitDataVolumeDir(dataVolume string) error {
	dataPath := strings.Split(dataVolume, ":")[0]
	return os.MkdirAll(dataPath, 0777)
}

func GetDefaultDataVolume() string {
	defaultDataPath := fmt.Sprintf("%s/tensorleap/data", getHomePath())
	return fmt.Sprintf("%s:%s", defaultDataPath, defaultDataPath)
}

func getHomePath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(fmt.Errorf("Failed to get home directory: %w", err))
	}

	return homeDir
}
