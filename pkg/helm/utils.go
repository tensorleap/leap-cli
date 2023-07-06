package helm

import (
	"log"
	"os"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/cli"
)

type Record = map[string]interface{}

const (
	RELEASE_NAME = "tensorleap"
	CHART_NAME   = "tensorleap"
	REPO_URL     = "https://helm.tensorleap.ai"
)

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
	ActionConfig *action.Configuration
	Settings *cli.EnvSettings
}

func CreateHelmConfig(kubeContext, namespace string) (*HelmConfig, error) {
	settings := cli.New()
	settings.SetNamespace(namespace)
	settings.KubeContext = kubeContext

	actionConfig := new(action.Configuration)
	if err := actionConfig.Init(settings.RESTClientGetter(), settings.Namespace(), os.Getenv("HELM_DRIVER"), log.Printf); err != nil {
		return nil, err
	}

	return &HelmConfig{
		Namespace:    namespace,
		ActionConfig: actionConfig,
		Settings: settings,
	}, nil
}
