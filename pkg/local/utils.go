package local

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"

	"github.com/tensorleap/cli-go/pkg/k3d"
	"github.com/tensorleap/cli-go/pkg/k8s"
	"github.com/tensorleap/cli-go/pkg/log"
)

const VAR_DIR = "/var/lib/tensorleap/standalone"
const KUBE_CONTEXT = "k3d-tensorleap"
const KUBE_NAMESPACE = "tensorleap"

func GetLatestImages(useGpu bool) (necessaryImages []string, backgroundImage string, err error) {
	resp, err := http.Get("https://raw.githubusercontent.com/tensorleap/helm-charts/master/images.txt")
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, "", fmt.Errorf("Getting latest chart images returned bad status code: %v", resp.StatusCode)
	}

	tensorleapImages, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "", err
	}
	k3sVersion := k3d.K3sVersion
	if useGpu {
		k3sVersion = k3d.K3sGpuVersion
	}

	resp, err = http.Get(fmt.Sprintf("https://github.com/k3s-io/k3s/releases/download/%s/k3s-images.txt", strings.Replace(k3sVersion, "-", "+", 1)))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		err = fmt.Errorf("Getting latest k3s images returned bad status code: %v", resp.StatusCode)
		return
	}

	k3sImages, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	allImages := strings.Split(string(tensorleapImages), "\n")
	allImages = append(allImages, strings.Split(string(k3sImages), "\n")...)

	necessaryImages = []string{}
	for _, img := range allImages {
		if strings.Contains(img, "engine") {
			backgroundImage = img
		} else if len(img) > 0 {
			necessaryImages = append(necessaryImages, img)
		}
	}

	return
}

func InitVarDir() error {
	_, err := os.Stat(VAR_DIR)
	if os.IsNotExist(err) {
		log.Printf("Creating directory: %s (you may be asked to enter the root user password)", VAR_DIR)
		mkdirCmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("sudo mkdir -p %s", VAR_DIR))
		if err := mkdirCmd.Run(); err != nil {
			return err
		}

		log.Println("Setting directory permissions")
		chmodCmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("sudo chmod -R 777 %s", VAR_DIR))
		if err := chmodCmd.Run(); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	return initVarDirSubDir()
}

func initVarDirSubDir() error {
	for _, dir := range []string{"storage", "registry", "logs"} {
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

// SetupInfra init VAR_DIR, setup VerboseLog and connect its output into a file
func SetupInfra(cmdName string) (closeLogFile func(), err error) {
	err = InitVarDir()
	if err != nil {
		return
	}

	k3d.SetupLogger(log.VerboseLogger)
	k8s.SetupLogger(log.VerboseLogger)

	logPath := createLogFilePath(cmdName)
	closeLogFile, err = log.ConnectFileToVerboseLogOutput(logPath)
	return
}

func createLogFilePath(cmdName string) string {
	filePath := fmt.Sprintf("%s/logs/%s_%s.log",
		VAR_DIR,
		cmdName,
		time.Now().Format("2006-01-02_15-04-05"),
	)
	return filePath
}
