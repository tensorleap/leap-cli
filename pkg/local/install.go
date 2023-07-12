package local

import (
	"fmt"
	"os"
	"strings"

	"github.com/tensorleap/cli-go/pkg/helm"
)

func InitDataVolumeDir(dataVolume string) error {
	dataPath := strings.Split(dataVolume, ":")[0]
	return os.MkdirAll(dataPath, 0777)
}

func GetDefaultDataVolume() string {
	defaultDataPath := fmt.Sprintf("%s/tensorleap/data", getHomePath())
	return fmt.Sprintf("%s:%s", defaultDataPath, defaultDataPath)
}

func InstallHelm(useGpu bool, dataContainerPath string) error {
	helmConfig, err := helm.CreateHelmConfig(KUBE_CONTEXT, KUBE_NAMESPACE)
	if err != nil {
		return err
	}

	isHelmReleaseExisted, err := helm.IsHelmReleaseExists(helmConfig)
	if err != nil {
		return err
	}
	if isHelmReleaseExisted {
		if err := helm.UpgradeTensorleapChartVersion(
			helmConfig,
		); err != nil {
			return err
		}
	} else {
		values := helm.CreateTensorleapChartValues(useGpu, dataContainerPath)
		if err := helm.InstallLatestTensorleapChartVersion(
			helmConfig,
			values,
		); err != nil {
			return err
		}
	}
	return nil
}

func getHomePath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(fmt.Errorf("Failed to get home directory: %w", err))
	}

	return homeDir
}
