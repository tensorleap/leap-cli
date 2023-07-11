package local

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/tensorleap/cli-go/pkg/k3d"
	"github.com/tensorleap/cli-go/pkg/k8s"
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

func SetupBackgroundLogger(logName string) *logrus.Logger {
	backgroundLogger := logrus.New()
	backgroundLogger.SetLevel(logrus.DebugLevel)
	backgroundLogger.SetOutput(io.Discard)
	k3d.SetupLogger(backgroundLogger)
	k8s.SetupLogger(backgroundLogger)
	return backgroundLogger
}
