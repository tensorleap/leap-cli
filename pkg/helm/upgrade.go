package helm

import (
	"time"

	"github.com/tensorleap/cli-go/pkg/log"
	"helm.sh/helm/v3/pkg/action"
)

func UpgradeTensorleapChartVersion(
	config *HelmConfig,
) error {

	log.Println("Upgrading helm chart (will take few minutes)")

	client := action.NewUpgrade(config.ActionConfig)
	client.ChartPathOptions.RepoURL = REPO_URL
	client.Namespace = config.Namespace
	client.Wait = true
	client.ReuseValues = true

	latestChart, err := getLatestChart(config, &client.ChartPathOptions)
	if err != nil {
		return err
	}

	client.Timeout = 20 * time.Minute

	_, err = client.RunWithContext(config.Context, RELEASE_NAME, latestChart, Record{})
	if err != nil {
		return err
	}

	log.Printf("Tensorleap upgrade on local k3d cluster! version: %s", latestChart.Metadata.Version)

	return nil
}
