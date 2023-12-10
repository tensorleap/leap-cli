package project

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/tensorleap/leap-cli/pkg/api"
	"github.com/tensorleap/leap-cli/pkg/auth"
	"github.com/tensorleap/leap-cli/pkg/entity"
	"github.com/tensorleap/leap-cli/pkg/hub"
	"github.com/tensorleap/leap-cli/pkg/log"
	"github.com/tensorleap/leap-cli/pkg/tensorleapapi"
)

func BuildProjectContext(ctx context.Context, projectEntity *ProjectEntity, schemaVersion int) (*hub.ProjectContext, error) {
	bgImageBlobUrl := ""
	if projectEntity.BgImagePath != nil {
		bgImageBlobUrl = fmt.Sprintf("projects/%s/%s", projectEntity.Cid, *projectEntity.BgImagePath)
	}
	urlRes, _, err := api.ApiClient.GetDownloadSignedUrl(ctx).GetDownloadSignedUrlParams(*tensorleapapi.NewGetDownloadSignedUrlParams(bgImageBlobUrl)).Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to get download signed url: %v", err)
	}
	res, err := http.Get(urlRes.GetUrl())
	if err != nil {
		return nil, fmt.Errorf("failed to download bg image: %v", err)
	}
	defer res.Body.Close()
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

func CopyProject(
	sourceCtx context.Context, sourceProject *ProjectEntity,
	targetCtx context.Context, targetProjectName string,
) error {

	sourceUrl, _ := api.GetAuthFromContext(sourceCtx)
	targetUrl, _ := api.GetAuthFromContext(targetCtx)

	fileAccess, err := getCopyPublishSignedUrl(sourceCtx, targetCtx, targetProjectName)
	if err != nil {
		return err
	}

	log.Infof("Copying project\n\tfrom: %s:%s\n\tto:   %s:%s", sourceProject.GetName(), sourceUrl, targetProjectName, targetUrl)

	err = PublishProject(sourceCtx, sourceProject.Cid, fileAccess)
	if err != nil {
		return err
	}

	targetProjectMeta := &hub.ProjectMeta{
		Name:            targetProjectName,
		Description:     sourceProject.GetDescription(),
		Tags:            sourceProject.Tags,
		Categories:      sourceProject.Categories,
		BgImagePath:     sourceProject.GetBgImagePath(),
		SourceProjectId: sourceProject.Cid,
	}

	err = ImportProject(targetCtx, targetProjectName, fileAccess.Get, targetProjectMeta)
	if err != nil {
		return err
	}
	return nil
}

func getCopyPublishSignedUrl(sourceCtx, targetCtx context.Context, fileName string) (*hub.FileAccessBySignedUrl, error) {
	sourceUrl, _ := api.GetAuthFromContext(sourceCtx)
	targetUrl, _ := api.GetAuthFromContext(targetCtx)

	isTargetLocal := auth.IsLocalUrl(targetUrl)
	isSourceLocal := auth.IsLocalUrl(sourceUrl)

	var ctx context.Context
	var origin *string
	if !isTargetLocal {
		ctx = targetCtx
	} else {
		ctx = sourceCtx
	}

	if isTargetLocal && isSourceLocal {
		// in case of local to local copy, we need to use the environment origin by setting origin to empty string
		emptyOrigin := ""
		origin = &emptyOrigin
	}
	signedUploadUrl, url, err := getTempUploadedSignedUrl(ctx, fileName, origin)
	if err != nil {
		return nil, fmt.Errorf("failed to get signed url for the uploaded project: %v", err)
	}
	signedGetUrl, err := getSignedUrl(ctx, url, http.MethodGet, time.Hour*24, origin)
	if err != nil {
		return nil, fmt.Errorf("failed to get singed url for the uploaded project: %v", err)
	}
	signedHeadUrl, err := getSignedUrl(ctx, url, http.MethodHead, time.Hour*24, origin)
	if err != nil {
		return nil, fmt.Errorf("failed to get singed url for the uploaded project: %v", err)
	}
	return &hub.FileAccessBySignedUrl{
		Put:  signedUploadUrl,
		Get:  signedGetUrl,
		Head: signedHeadUrl,
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
