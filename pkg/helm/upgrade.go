package helm

import (
	"context"
	"log"

	"helm.sh/helm/v3/pkg/action"
)

func UpgradeTensorleapChartVersion(
	ctx context.Context,
	config *HelmConfig,
) error {

	client := action.NewUpgrade(config.ActionConfig)
	client.ChartPathOptions.RepoURL = REPO_URL
	client.Namespace = config.Namespace
	client.Wait = true
	client.ReuseValues = true

	latestChart, err := getLatestChart(config, &client.ChartPathOptions)
	if err != nil {
		return err
	}

	_, err = client.RunWithContext(ctx, RELEASE_NAME, latestChart, Record{})
	if err != nil {
		return err
	}

	log.Printf("Tensorleap upgrade on local k3d cluster! version: %s", latestChart.Metadata.Version)

	return nil
}
