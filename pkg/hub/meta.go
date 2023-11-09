package hub

import (
	"fmt"

	"github.com/tensorleap/leap-cli/pkg/log"
	"gopkg.in/yaml.v3"
)

type ProjectMeta struct {
	Name            string                 `json:"name" yaml:"name"`
	Description     string                 `json:"description,omitempty" yaml:"description,omitempty"`
	Tags            []string               `json:"tags" yaml:"tags"`
	Categories      map[string]interface{} `json:"categories" yaml:"categories"`
	BgImagePath     string                 `json:"bgImagePath" yaml:"bgImagePath"`
	SchemaVersion   int                    `json:"schemaVersion" yaml:"schemaVersion"`
	SourceProjectId string                 `json:"sourceProjectId" yaml:"sourceProjectId"`
}

type HubProjectVersion struct {
	ProjectMeta `yaml:",inline"`
	CreatedAt   int64  `json:"createdAt" yaml:"createdAt"`
	Path        string `json:"path" yaml:"path"`
	Size        int    `json:"size" yaml:"size"` // bytes
}

type HubProject struct {
	CreatedAt int64                        `json:"createdAt" yaml:"createdAt"`
	Versions  map[string]HubProjectVersion `json:"versions" yaml:"versions"`
}

type HubMeta struct {
	Namespace string                `json:"namespace" yaml:"namespace"`
	Projects  map[string]HubProject `json:"projects" yaml:"projects"`
}

type HubMetaWrapper struct {
	Data HubMeta `json:"data" yaml:"data"`
}

func CreateWrapperMetaFromUrl(url string) (*HubMetaWrapper, error) {
	metaUrl := fmt.Sprintf("%s/%s", url, NAMESPACE_META_FILE_NAME)
	metaBytes, err := downloadFile(metaUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to download meta file: %v", err)
	}
	meta, err := UnmarshalHubMeta(metaBytes)
	if err != nil {
		return nil, err
	}
	return NewHubMetaWrapper(*meta), nil
}

func UnmarshalHubMeta(data []byte) (*HubMeta, error) {
	var meta HubMeta
	err := yaml.Unmarshal(data, &meta)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal meta: %v", err)
	}

	return &meta, nil
}

func NewHubMetaWrapper(hubMeta HubMeta) *HubMetaWrapper {
	return &HubMetaWrapper{
		Data: hubMeta,
	}
}

func (h *HubMetaWrapper) GetVersionMeta(projectName string, projectVersion int) (*HubProjectVersion, error) {
	projectMeta, ok := h.Data.Projects[projectName]
	if !ok {
		return nil, fmt.Errorf("project '%s' not found", projectName)
	}
	versionStr := fmt.Sprint(projectVersion)
	versionMeta, found := projectMeta.Versions[versionStr]
	if !found {
		return nil, fmt.Errorf("project '%s' version '%d' not found", projectName, projectVersion)
	}
	return &versionMeta, nil
}

func (h *HubMetaWrapper) GetProjectNames() []string {
	keys := make([]string, 0, len(h.Data.Projects))
	for k := range h.Data.Projects {
		keys = append(keys, k)
	}
	return keys
}

func (h *HubMetaWrapper) GetProjectNamesByVersion(schemeVersion int) []string {
	keys := make([]string, 0, len(h.Data.Projects))
	schemeVersionStr := fmt.Sprint(schemeVersion)
	for k := range h.Data.Projects {
		_, isVersionExists := h.Data.Projects[k].Versions[schemeVersionStr]
		if isVersionExists {
			keys = append(keys, k)
		}
	}
	return keys
}

func (h *HubMetaWrapper) GetProjectVersions(projectName string) []int {
	versionMap := h.Data.Projects[projectName].Versions
	versions := make([]int, 0, len(versionMap))
	for _, v := range versionMap {
		versions = append(versions, v.SchemaVersion)
	}
	return versions
}

func (h *HubMetaWrapper) ProjectVersionExists(projectName string, projectVersion int) bool {
	for _, v := range h.Data.Projects[projectName].Versions {
		if v.SchemaVersion == projectVersion {
			return true
		}
	}
	return false
}

func (h *HubMetaWrapper) UpdateVersions(projectVersions []HubProjectVersion) {
	for _, projectVersion := range projectVersions {
		log.Printf("Updating meta - add project: '%s' version: %d\n", projectVersion.Name, projectVersion.SchemaVersion)

		projectMeta, ok := h.Data.Projects[projectVersion.Name]
		if !ok {
			projectMeta = HubProject{
				CreatedAt: projectVersion.CreatedAt,
				Versions:  map[string]HubProjectVersion{},
			}
			h.Data.Projects[projectVersion.Name] = projectMeta
		}
		projectMeta.Versions[fmt.Sprint(projectVersion.SchemaVersion)] = projectVersion
	}
}

func (h *HubMetaWrapper) DeleteProject(projectName string) bool {
	if _, ok := h.Data.Projects[projectName]; ok {
		log.Printf("Updating meta - delete project: '%s'\n", projectName)
		delete(h.Data.Projects, projectName)
		return true
	}
	return false
}

func (h *HubMetaWrapper) DeleteProjectVersion(projectName string, projectVersion int) bool {
	projectMeta, ok := h.Data.Projects[projectName]
	if ok {
		for _, v := range projectMeta.Versions {
			if v.SchemaVersion == projectVersion {
				delete(projectMeta.Versions, fmt.Sprint(projectVersion))
				log.Printf("Updating meta - delete project: '%s' version: %d\n", projectName, projectVersion)
				return true
			}
		}
	}
	return false
}
