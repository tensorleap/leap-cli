package datasets

import (
	"context"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	. "github.com/tensorleap/cli-go/pkg/api"
	"github.com/tensorleap/cli-go/pkg/local"
	"github.com/tensorleap/cli-go/pkg/log"
	"github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

var EmptyDatasetVersionError = fmt.Errorf("Dataset is empty")

func AskForDatasetName() (name string, err error) {
	prompt := &survey.Input{
		Message: "Enter dataset name",
	}
	err = survey.AskOne(prompt, &name)
	return
}

func AskForDataset(ctx context.Context) (*tensorleapapi.Dataset, error) {
	data, _, err := ApiClient.GetDatasets(ctx).Execute()
	if err != nil {
		return nil, err
	}
	if len(data.Datasets) == 0 {
		return nil, fmt.Errorf("No datasets found, please create dataset")
	}
	index := -1
	options := []string{}
	for _, dataset := range data.Datasets {
		options = append(options, fmt.Sprintf("%s (%s)", dataset.GetName(), dataset.GetCid()))
	}
	prompt := &survey.Select{
		Message: "Choose dataset:",
		Options: options,
	}
	err = survey.AskOne(prompt, &index)
	if err != nil || index == -1 {
		return nil, fmt.Errorf("No dataset selected")
	}
	return &data.Datasets[index], nil
}

func CreateNewDataset(ctx context.Context, name string) (*tensorleapapi.Dataset, error) {
	log.Println("Creating dataset:", name)
	dataset, _, err := ApiClient.AddDataset(ctx).
		NewDatasetParams(*tensorleapapi.NewNewDatasetParams(*tensorleapapi.NewNullableString(&name))).
		Execute()

	if err != nil {
		return nil, fmt.Errorf("Failed to create a new dataset: %v", err)
	}
	return &dataset.Dataset, nil
}

func GetDatasetById(ctx context.Context, datasetId string) (*tensorleapapi.Dataset, error) {
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

func GetLatestVersion(ctx context.Context, datasetId string) (*tensorleapapi.DatasetVersion, error) {
	version, _, err := ApiClient.GetLatestDatasetVersion(ctx).
		GetLatestDatasetVersionParams(*tensorleapapi.NewGetLatestDatasetVersionParams(datasetId)).
		Execute()
	if err != nil {
		return nil, fmt.Errorf("Failed to get latest version for datasetId: %s", datasetId)
	}
	if isDatasetVersionEmpty(version.LatestVersion) {
		return version.LatestVersion, EmptyDatasetVersionError
	}
	return version.LatestVersion, nil
}

func isDatasetVersionEmpty(datasetVersion *tensorleapapi.DatasetVersion) bool {
	return len(datasetVersion.GetBlobPath()) == 0
}

func CloneDatasetVersionCode(ctx context.Context, datasetVersion *tensorleapapi.DatasetVersion, outputDir string) ([]string, error) {
	if isDatasetVersionEmpty(datasetVersion) {
		return []string{}, EmptyDatasetVersionError
	}
	blobPath := datasetVersion.GetBlobPath()
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
