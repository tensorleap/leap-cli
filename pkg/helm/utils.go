package helm

import (
	"context"
	"log"
	"os"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/storage/driver"
)

type Record = map[string]interface{}

const (
	RELEASE_NAME = "tensorleap"
	CHART_NAME   = "tensorleap"
	REPO_URL     = "https://helm.tensorleap.ai"
)

func IsHelmReleaseExists(config *HelmConfig) (bool, error) {
	client := action.NewHistory(config.ActionConfig)
	client.Max = 1
	_, err := client.Run(RELEASE_NAME)
	if err == driver.ErrReleaseNotFound {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}

func CreateTensorleapChartValues(useGpu bool, dataDir string) Record {
	return Record{
		"tensorleap-engine": Record{
			"gpu":                useGpu,
			"localDataDirectory": dataDir,
		},
	}
}

type HelmConfig struct {
	Namespace    string
	Context context.Context
	ActionConfig *action.Configuration
	Settings     *cli.EnvSettings
}

func CreateHelmConfig(kubeContext, namespace string) (*HelmConfig, error) {
	settings := cli.New()
	settings.SetNamespace(namespace)
	settings.KubeContext = kubeContext

	// Any other context with cancel will failed immediately when running helm actions, using background context solve it
	ctx := context.Background()

	actionConfig := new(action.Configuration)
	if err := actionConfig.Init(settings.RESTClientGetter(), settings.Namespace(), os.Getenv("HELM_DRIVER"), log.Printf); err != nil {
		return nil, err
	}

	return &HelmConfig{
		Context: ctx,
		Namespace:    namespace,
		ActionConfig: actionConfig,
		Settings:     settings,
	}, nil
}

func getLatestChart(config *HelmConfig, chartPathOptions *action.ChartPathOptions) (*chart.Chart, error) {
	chartPath, err := chartPathOptions.LocateChart(CHART_NAME, config.Settings)
	if err != nil {
		return nil, err
	}

	chartRequested, err := loader.Load(chartPath)
	if err != nil {
		return nil, err
	}
	return chartRequested, nil
}
