package helm

import (
	"context"
	"log"
	"os"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
)

func InstallLatestTensorleapChartVersion(ctx context.Context, kubeContext string, namespace string) error {
	settings := cli.New()
	settings.SetNamespace(namespace)
	settings.KubeContext = kubeContext

	actionConfig := new(action.Configuration)
	if err := actionConfig.Init(settings.RESTClientGetter(), settings.Namespace(), os.Getenv("HELM_DRIVER"), log.Printf); err != nil {
		log.Printf("%+v", err)
		os.Exit(1)
	}

	client := action.NewInstall(actionConfig)
	client.ChartPathOptions.RepoURL = "https://helm.tensorleap.ai"
	client.Namespace = settings.Namespace()
	client.CreateNamespace = true
	client.Wait = true
	client.ReleaseName = "tensorleap"

	// client := action.NewPull()
	cp, err := client.ChartPathOptions.LocateChart("tensorleap", settings)
	if err != nil {
		return err
	}

	chartRequested, err := loader.Load(cp)
	if err != nil {
		return err
	}

	vals := make(map[string]interface{})

	_, err = client.RunWithContext(ctx, chartRequested, vals)
	if err != nil {
		return err
	}

	log.Printf("Tensorleap installed on local k3d cluster! version: %s", chartRequested.Metadata.Version)

	return nil
}
