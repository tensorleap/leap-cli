package local

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/tensorleap/cli-go/pkg/k3d"
)

const VAR_DIR = "/var/lib/tensorleap/standalone"

func GetLatestImages(useGpu bool) ([]string, string, error) {
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
		return nil, "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, "", fmt.Errorf("Getting latest k3s images returned bad status code: %v", resp.StatusCode)
	}

	k3sImages, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "", err
	}

	allImages := strings.Split(string(tensorleapImages), "\n")
	allImages = append(allImages, strings.Split(string(k3sImages), "\n")...)

	var engineImage string
	ret := []string{}
	for _, img := range allImages {
		if len(img) > 0 {
			if strings.Contains(img, "engine") {
				engineImage = img
			} else {
				ret = append(ret, img)
			}
		}
	}

	return ret, engineImage, nil
}