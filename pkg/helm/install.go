package helm

import (
	"context"
	"log"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
)

func InstallLatestTensorleapChartVersion(
	ctx context.Context,
	config *HelmConfig,
	values Record,
) error {
	

	client := action.NewInstall(config.ActionConfig)
	client.ChartPathOptions.RepoURL = REPO_URL
	client.Namespace = config.Namespace
	client.CreateNamespace = true
	client.Wait = true
	client.ReleaseName = RELEASE_NAME

	// client := action.NewPull()
	cp, err := client.ChartPathOptions.LocateChart(CHART_NAME, config.Settings)
	if err != nil {
		return err
	}

	chartRequested, err := loader.Load(cp)
	if err != nil {
		return err
	}

	_, err = client.RunWithContext(ctx, chartRequested, values)
	if err != nil {
		return err
	}

	log.Printf("Tensorleap installed on local k3d cluster! version: %s", chartRequested.Metadata.Version)

	return nil
}

