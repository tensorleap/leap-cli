package hub

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
	"github.com/tensorleap/leap-cli/pkg/log"
	"k8s.io/kubectl/pkg/util/slice"
)

func ValidateOrSelectProject(projectFromFlag string, projectNames []string) (string, error) {
	if projectFromFlag != "" {
		exists := slice.ContainsString(projectNames, projectFromFlag, nil)
		if exists {
			return projectFromFlag, nil
		}
		log.Warnf("Project %s not found", projectFromFlag)
	}

	prompt := &survey.Select{
		Message: "Select project",
		Options: projectNames,
	}
	index := -1
	err := survey.AskOne(prompt, &index)
	if err != nil || index == -1 {
		return "", fmt.Errorf("no project selected")
	}
	return projectNames[index], nil
}

func ValidateOrSelectLatestVersion(projectName string, versionFlag string, versions []int) (int, error) {
	if versionFlag != "" {
		versionFlag, _ := strconv.Atoi(versionFlag)
		exists := contains(versions, versionFlag)
		if exists {
			return versionFlag, nil
		}
		return 0, fmt.Errorf("not found version %v, for project %s", versionFlag, projectName)
	}
	latestVersion := 0

	for _, version := range versions {
		if version > latestVersion {
			latestVersion = version
		}
	}

	return latestVersion, nil
}

func GetPublicUrl() string {
	return fmt.Sprintf("%s/%s", PUBLIC_STORAGE_HUB_URL, DEFAULT_NAMESPACE)
}

func downloadFile(url string) ([]byte, error) {
	res, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("failed to download file: %v", err)
	}
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}
	return b, nil
}

func contains(arr []int, target int) bool {
	for _, num := range arr {
		if num == target {
			return true
		}
	}
	return false
}

var ErrProjectExists = fmt.Errorf("project already exists, use -o, --override to re-upload")

func HandleProjectExistsOnPublish(ctx context.Context, hubApi *HubApi, meta *HubMetaWrapper, projectMeta *ProjectMeta, override bool) error {
	projectName := projectMeta.Name
	projectVersion := projectMeta.SchemaVersion

	isProjectExists := meta.ProjectVersionExists(
		projectName,
		projectVersion,
	)
	if isProjectExists {
		if !override {
			return ErrProjectExists
		}
		log.Infof("Removing previous version")
		err := hubApi.DeleteProjectVersion(projectName, projectVersion)
		if err != nil {
			return err
		}
	}
	return nil
}

func IsHubEnabled() bool {
	return os.Getenv("LEAP_HUB_ENABLED") == "true"
}
