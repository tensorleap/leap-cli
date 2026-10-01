package project

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/tensorleap/leap-cli/pkg/api"
	"github.com/tensorleap/leap-cli/pkg/entity"
	"github.com/tensorleap/leap-cli/pkg/hub"
)

func BuildProjectContext(ctx context.Context, projectEntity *ProjectEntity, schemaVersion int) (*hub.ProjectContext, error) {
	if projectEntity.BgImagePath == nil || *projectEntity.BgImagePath == "" {
		return nil, fmt.Errorf("project %s has no background image configured", projectEntity.Name)
	}

	bgImageBlobUrl := fmt.Sprintf("projects/%s/%s", projectEntity.Cid, *projectEntity.BgImagePath)
	downloadUrl, err := api.GetDownloadSignedUrl(ctx, bgImageBlobUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to get download signed url: %v", err)
	}
	res, err := http.Get(downloadUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to download bg image: %v", err)
	}
	defer func() { _ = res.Body.Close() }()
	bgImageBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read bg image: %v", err)
	}

	return &hub.ProjectContext{
		Meta: hub.ProjectMeta{
			Name:            projectEntity.Name,
			SchemaVersion:   schemaVersion,
			BgImagePath:     projectEntity.GetBgImagePath(),
			Description:     projectEntity.GetDescription(),
			Tags:            projectEntity.Tags,
			SourceProjectId: projectEntity.Cid,
			Categories:      projectEntity.Categories,
		},
		BgImage: hub.Image{
			Name:   *projectEntity.BgImagePath,
			Buffer: bgImageBytes,
		},
	}, nil
}

func ValidateProjectName(projectName, defaultProjectName string, projects []ProjectEntity) (string, error) {
	existedNames := entity.GetNames(projects, ProjectEntityDesc)
	err := entity.CreateUniqueNameValidator(existedNames)(projectName)
	if err == nil {
		return projectName, nil
	}
	if defaultProjectName == "" {
		defaultProjectName = projectName
	}
	return entity.AskForName(existedNames, defaultProjectName, ProjectEntityDesc)
}

func extractUrl(rowUrl string) string {
	parsedURL, _ := url.Parse(rowUrl)
	return parsedURL.Scheme + "://" + parsedURL.Host + parsedURL.Path
}
