package code

import (
	"context"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	. "github.com/tensorleap/leap-cli/pkg/api"
	"github.com/tensorleap/leap-cli/pkg/local"
	"github.com/tensorleap/leap-cli/pkg/log"
	"github.com/tensorleap/leap-cli/pkg/tensorleapapi"
)

var ErrEmptyCodeIntegrationVersion = fmt.Errorf("CodeIntegration is empty")

func AskForCodeIntegrationName() (name string, err error) {
	prompt := &survey.Input{
		Message: "Enter dataset name",
	}
	err = survey.AskOne(prompt, &name)
	return
}

func AskForCodeIntegration(ctx context.Context) (*tensorleapapi.Dataset, error) {
	data, _, err := ApiClient.GetDatasets(ctx).Execute()
	if err != nil {
		return nil, err
	}
	if len(data.Datasets) == 0 {
		return nil, fmt.Errorf("No code integration found, please create dataset")
	}
	index := -1
	options := []string{}
	for _, dataset := range data.Datasets {
		options = append(options, fmt.Sprintf("%s (%s)", dataset.GetName(), dataset.GetCid()))
	}
	prompt := &survey.Select{
		Message: "Choose code integration:",
		Options: options,
	}
	err = survey.AskOne(prompt, &index)
	if err != nil || index == -1 {
		return nil, fmt.Errorf("No code integration selected")
	}
	return &data.Datasets[index], nil
}

func CreateNewCodeIntegration(ctx context.Context, name string) (*tensorleapapi.Dataset, error) {
	log.Println("Creating dataset:", name)
	dataset, _, err := ApiClient.AddDataset(ctx).
		NewDatasetParams(*tensorleapapi.NewNewDatasetParams(*tensorleapapi.NewNullableString(&name))).
		Execute()

	if err != nil {
		return nil, fmt.Errorf("Failed to create a new code integration: %v", err)
	}
	return &dataset.Dataset, nil
}

func GetCodeIntegrationById(ctx context.Context, datasetId string) (*tensorleapapi.Dataset, error) {
	data, _, err := ApiClient.GetDatasets(ctx).Execute()
	if err != nil {
		return nil, err
	}
	var selectedDataset *tensorleapapi.Dataset = nil
	for _, dataset := range data.Datasets {
		if dataset.GetCid() == datasetId {
			fmt.Println("Found dataset:", dataset.GetName())
			selectedDataset = &dataset
			break
		}
	}
	if selectedDataset == nil {
		return nil, fmt.Errorf("Didn't find dataset with id: %v", datasetId)
	}
	return selectedDataset, nil
}

func GetLatestVersion(ctx context.Context, codeIntegrationId string) (*tensorleapapi.DatasetVersion, error) {
	version, _, err := ApiClient.GetLatestDatasetVersion(ctx).
		GetLatestDatasetVersionParams(*tensorleapapi.NewGetLatestDatasetVersionParams(codeIntegrationId)).
		Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to get latest version for code integration id: %s", codeIntegrationId)
	}
	if isDatasetVersionEmpty(version.LatestVersion) {
		return version.LatestVersion, ErrEmptyCodeIntegrationVersion
	}
	return version.LatestVersion, nil
}

func isDatasetVersionEmpty(datasetVersion *tensorleapapi.DatasetVersion) bool {
	return len(datasetVersion.GetBlobPath()) == 0
}

func CloneCodeIntegrationVersion(ctx context.Context, codeIntegrationVersion *tensorleapapi.DatasetVersion, outputDir string) ([]string, error) {
	if isDatasetVersionEmpty(codeIntegrationVersion) {
		return []string{}, ErrEmptyCodeIntegrationVersion
	}
	blobPath := codeIntegrationVersion.GetBlobPath()
	res, _, err := ApiClient.GetDownloadSignedUrl(ctx).
		GetDownloadSignedUrlParams(*tensorleapapi.NewGetDownloadSignedUrlParams(blobPath)).
		Execute()
	if err != nil {
		return nil, err
	}

	files, err := local.DownloadAndExtractTarFile(res.GetUrl(), outputDir)
	if err != nil {
		return nil, err
	}
	return files, nil
}
