package server

import (
	"fmt"
	"os"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/tensorleap/leap-cli/pkg/helm"
	"github.com/tensorleap/leap-cli/pkg/log"
)

func InitUseGPU(useGpu *bool, useCpu bool) error {
	if *useGpu || useCpu {
		return nil
	}

	prompt := survey.Confirm{
		Message: "Do you want to use GPU?",
		Default: false,
	}

	err := survey.AskOne(&prompt, useGpu)
	return err
}

func InitDatasetDirectory(datasetDirectory *string) error {
	defaultDatasetDirectory := GetDefaultDataVolume()

	if *datasetDirectory == "" {
		fromPath := ""
		prompt := survey.Input{
			Message: "Enter dataset directory:",
			Default: defaultDatasetDirectory,
		}
		err := survey.AskOne(&prompt, &fromPath)
		if err != nil {
			return err
		}
		*datasetDirectory = fromPath
	}
	if !strings.Contains(*datasetDirectory, ":") {
		toPath := ""
		prompt := survey.Input{
			Message: "Enter container dataset directory:",
			Default: *datasetDirectory,
		}
		err := survey.AskOne(&prompt, &toPath)
		if err != nil {
			return err
		}
		*datasetDirectory = fmt.Sprintf("%s:%s", *datasetDirectory, toPath)
	}
	log.SendCloudReport("info", "Init data volume", "Starting",
		&map[string]interface{}{"params": map[string]interface{}{"datasetDirectory": datasetDirectory}},
	)

	dataPath := strings.Split(*datasetDirectory, ":")[0]
	return os.MkdirAll(dataPath, 0777)
}

func GetDefaultDataVolume() string {
	defaultDataPath := fmt.Sprintf("%s/tensorleap/data", getHomePath())
	return defaultDataPath
}

func InstallHelm(useGpu bool, dataContainerPath string, disableMetrics bool) error {
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
		values := helm.CreateTensorleapChartValues(useGpu, dataContainerPath, disableMetrics)
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
		panic(fmt.Errorf("failed to get home directory: %w", err))
	}

	return homeDir
}
