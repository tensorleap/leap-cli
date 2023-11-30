package hub

import "io"

type HubApi struct {
	hubStorage *HubStorageApi
}

func NewHubApi(hubStorage *HubStorageApi) *HubApi {
	return &HubApi{hubStorage: hubStorage}
}

func (h *HubApi) PublishProjectContentBySignedUrl(meta *ProjectMeta) (*FileAccessBySignedUrl, *ProjectFilePaths, error) {
	_, err := h.hubStorage.DeleteProjectVersion(meta.Name, meta.SchemaVersion)
	if err != nil {
		return nil, nil, err
	}

	return h.hubStorage.UploadProjectContentBySignedUrl(meta)
}

func (h *HubApi) PublishProject(file io.Reader, ctx *ProjectContext) error {
	_, err := h.hubStorage.DeleteProjectVersion(ctx.Meta.Name, ctx.Meta.SchemaVersion)
	if err != nil {
		return err
	}

	files, err := h.hubStorage.UploadProjectContent(file, &ctx.Meta)
	if err != nil {
		return err
	}
	return h.PublishProjectMeta(files, ctx)
}

func (h *HubApi) PublishProjectMeta(files *ProjectFilePaths, projectCtx *ProjectContext) error {
	projectMetaVersion, err := h.hubStorage.UpdateProjectMeta(files, projectCtx)
	if err != nil {
		return err
	}

	hubMeta, err := h.hubStorage.GetMetaWrapper()
	if err != nil {
		return err
	}

	hubMeta.UpdateVersions([]HubProjectVersion{*projectMetaVersion})
	return h.hubStorage.SetMeta(hubMeta.Data)
}

func (h *HubApi) DownloadProject(projectName string, version int, destFileName string) error {
	return h.hubStorage.DownloadProject(projectName, version, destFileName)
}

func (h *HubApi) DeleteProject(projectName string) error {
	_, err := h.hubStorage.DeleteProject(projectName)
	if err != nil {
		return err
	}

	meta, err := h.hubStorage.GetMetaWrapper()
	if err != nil {
		return err
	}

	metaUpdated := meta.DeleteProject(projectName)
	if metaUpdated {
		return h.hubStorage.SetMeta(meta.Data)
	}

	return nil
}

func (h *HubApi) DeleteProjectVersion(projectName string, projectVersion int) error {
	_, err := h.hubStorage.DeleteProjectVersion(projectName, projectVersion)
	if err != nil {
		return err
	}

	meta, err := h.hubStorage.GetMetaWrapper()
	if err != nil {
		return err
	}

	metaUpdated := meta.DeleteProjectVersion(projectName, projectVersion)
	if metaUpdated {
		return h.hubStorage.SetMeta(meta.Data)
	}

	return nil
}
