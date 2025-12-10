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

func CopyProject(
	sourceCtx context.Context, sourceProject *ProjectEntity,
	targetCtx context.Context, targetProjectName string,
	exportOptions ExportProjectParams,
	waitForImport bool,
) error {

	sourceUrl, _ := api.GetAuthFromContext(sourceCtx)
	targetUrl, _ := api.GetAuthFromContext(targetCtx)

	copyTo, err := getCopyToSignedUrl(sourceCtx, targetCtx, targetProjectName)
	if err != nil {
		return err
	}
	var copyToUrl string
	if copyTo != nil {
		copyToUrl = copyTo.Put
	}

	log.Infof("Copying project\n\tfrom: %s:%s\n\tto:   %s:%s", sourceProject.GetName(), sourceUrl, targetProjectName, targetUrl)

	exportJob, err := ExportProject(sourceCtx, sourceProject.Cid, copyToUrl, exportOptions)
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

	var copyFromUrl string
	if copyTo != nil {
		copyFromUrl = copyTo.Get
	} else {
		exportUrl := exportJob.Params.ExportProjectParams.GetExportUrl()
		var origin *string
		isCopyFromLocalToLocal := auth.IsLocalUrl(sourceUrl)
		if isCopyFromLocalToLocal {
			emptyOriginUseServerStorageUrl := ""
			origin = &emptyOriginUseServerStorageUrl
		}

		copyFromUrl, err = api.GetSignedUrl(sourceCtx, exportUrl, http.MethodGet, time.Hour*24, origin)
		if err != nil {
			return fmt.Errorf("failed to get signed url for exported project file : %v", err)
		}
	}

	err = ImportProject(targetCtx, targetProjectName, copyFromUrl, targetProjectMeta, waitForImport)
	if err != nil {
		return err
	}
	return nil
}

func getCopyToSignedUrl(sourceCtx, targetCtx context.Context, fileName string) (*hub.FileAccessBySignedUrl, error) {
	sourceUrl, _ := api.GetAuthFromContext(sourceCtx)
	targetUrl, _ := api.GetAuthFromContext(targetCtx)

	isSourceLocal := auth.IsLocalUrl(sourceUrl)

	isSameEnv := sourceUrl == targetUrl

	if !isSourceLocal || isSameEnv {
		return nil, nil
	}

	ctx := targetCtx

	res, err := api.GetUploadSignedUrl(ctx, fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to get signed url for the uploaded project: %v", err)
	}
	signedUploadUrl := res.GetUrl()
	url := res.GetFileName()

	signedGetUrl, err := api.GetSignedUrl(ctx, url, http.MethodGet, time.Hour*24, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get singed url for the uploaded project: %v", err)
	}

	return &hub.FileAccessBySignedUrl{
		Put: signedUploadUrl,
		Get: signedGetUrl,
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
