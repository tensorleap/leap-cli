package server

import (
	"fmt"
	"os"
	"strings"

	"github.com/tensorleap/leap-cli/pkg/helm"
	"github.com/tensorleap/leap-cli/pkg/log"
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
	log.SendCloudReport("info", "Installing helm", "Running", nil)
	helmConfig, err := helm.CreateHelmConfig(KUBE_CONTEXT, KUBE_NAMESPACE)
	if err != nil {
		log.SendCloudReport("error", "Failed creating helm config", "Failed",
			&map[string]interface{}{"kubeContext": KUBE_CONTEXT, "kubeNamespace": KUBE_NAMESPACE, "error": err.Error()})
		return err
	}

	isHelmReleaseExisted, err := helm.IsHelmReleaseExists(helmConfig)
	if err != nil {
		log.SendCloudReport("error", "Failed checking if helm release exists", "Failed",
			&map[string]interface{}{"helmConfig": helmConfig, "error": err.Error()})
		return err
	}
	if isHelmReleaseExisted {
		log.SendCloudReport("info", "Running helm upgrade", "Running", &map[string]interface{}{"isHelmReleaseExisted": isHelmReleaseExisted})
		if err := helm.UpgradeTensorleapChartVersion(
			helmConfig,
		); err != nil {
			log.SendCloudReport("error", "Failed upgrading helm charts versions", "Failed",
				&map[string]interface{}{"isHelmReleaseExisted": isHelmReleaseExisted, "helmConfig": helmConfig, "error": err.Error()})
			return err
		}
	} else {
		values := helm.CreateTensorleapChartValues(useGpu, dataContainerPath)
		log.SendCloudReport("info", "Setting up helm repo", "Running", &map[string]interface{}{"isHelmReleaseExisted": isHelmReleaseExisted})
		if err := helm.InstallLatestTensorleapChartVersion(
			helmConfig,
			values,
		); err != nil {
			log.SendCloudReport("error", "Failed installing latest chart versions", "Failed",
				&map[string]interface{}{"isHelmReleaseExisted": isHelmReleaseExisted, "helmConfig": helmConfig, "values": values, "error": err.Error()})
			return err
		}
	}

	log.SendCloudReport("info", "Successfully installed helm charts", "Running", nil)
	return nil
}

func getHomePath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(fmt.Errorf("Failed to get home directory: %w", err))
	}

	return homeDir
}
