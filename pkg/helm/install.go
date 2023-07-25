package helm

import (
	"time"

	"github.com/tensorleap/leap-cli/pkg/log"
	"helm.sh/helm/v3/pkg/action"
)

func InstallLatestTensorleapChartVersion(
	config *HelmConfig,
	values Record,
) error {

	log.Println("Installing helm chart (will take few minutes)")

	client := action.NewInstall(config.ActionConfig)
	client.ChartPathOptions.RepoURL = REPO_URL
	client.Namespace = config.Namespace
	client.CreateNamespace = true
	client.Wait = true
	client.ReleaseName = RELEASE_NAME

	latestChart, err := getLatestChart(config, &client.ChartPathOptions)
	if err != nil {
		return err
	}

	client.Timeout = 20 * time.Minute

	_, err = client.RunWithContext(config.Context, latestChart, values)
	if err != nil {
		return err
	}

	log.Printf("Tensorleap installed on local k3d cluster! version: %s", latestChart.Metadata.Version)

	return nil
}
