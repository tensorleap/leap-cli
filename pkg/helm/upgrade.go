package helm

import (
	"fmt"
	"time"

	"github.com/tensorleap/leap-cli/pkg/log"
	"helm.sh/helm/v3/pkg/action"
)

func UpgradeTensorleapChartVersion(
	config *HelmConfig,
	values Record,
) error {

	log.SendCloudReport("info", "Upgrading helm chart", "Running", nil)
	log.Println("Upgrading helm chart (will take few minutes)")

	client := action.NewUpgrade(config.ActionConfig)
	client.ChartPathOptions.RepoURL = REPO_URL
	client.Namespace = config.Namespace
	client.Wait = true
	if values == nil {
		oldValues, err := GetValues(config, RELEASE_NAME)
		if err != nil {
			return err
		}
		values, err = CreateTensorleapChartValuesFormOldValues(oldValues)
		if err != nil {
			return fmt.Errorf("failed creating values from previous helm release: %w", err)
		}
	}

	latestChart, err := getLatestChart(config, &client.ChartPathOptions)
	if err != nil {
		return err
	}

	client.Timeout = 20 * time.Minute

	_, err = client.RunWithContext(config.Context, RELEASE_NAME, latestChart, values)
	if err != nil {
		log.SendCloudReport("error", "Failed upgrading helm chart", "Failed",
			&map[string]interface{}{"releaseName": RELEASE_NAME, "latestChart": latestChart, "error": err.Error()})
		return err
	}

	log.Printf("Tensorleap upgrade on local k3d cluster! version: %s", latestChart.Metadata.Version)
	log.SendCloudReport("info", "Successfully upgraded helm chart", "Running", nil)

	return nil
}
