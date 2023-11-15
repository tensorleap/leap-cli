package project

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/tensorleap/leap-cli/pkg/api"
	"github.com/tensorleap/leap-cli/pkg/hub"
	"github.com/tensorleap/leap-cli/pkg/tensorleapapi"
)

func BuildProjectContext(ctx context.Context, projectEntity *ProjectEntity, schemaVersion int) (*hub.ProjectContext, error) {
	bgImageBlobUrl := fmt.Sprintf("projects/%s/%s", projectEntity.Cid, *projectEntity.BgImagePath)
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
