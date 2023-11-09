package hub

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetProjectNames(t *testing.T) {
	meta := loadMetaWrapper(t)

	projectNames := meta.GetProjectNames()
	assert.Equal(t, 12, len(projectNames))

	projectNamesByVersion := meta.GetProjectNamesByVersion(90)
	assert.Equal(t, 4, len(projectNamesByVersion))

	projectName := "IMDB"
	versions := meta.GetProjectVersions(projectName)
	assert.Equal(t, 1, len(versions))
}

func TestDeleteProject(t *testing.T) {
	meta := loadMetaWrapper(t)

	projectNames := meta.GetProjectNames()
	assert.Equal(t, 12, len(projectNames))

	projectName := "IMDB"
	meta.DeleteProject(projectName)

	projectNames = meta.GetProjectNames()
	assert.Equal(t, 11, len(projectNames))
	assert.NotContains(t, projectNames, projectName)
}

func TestCreateWrapperMetaFromPublicUrl(t *testing.T) {
	t.Skip("skipping test, because it requires internet connection")
	url := GetPublicUrl()
	meta, err := CreateWrapperMetaFromUrl(url)
	if err != nil {
		t.Errorf("failed to create wrapper meta from public url: %v", err)
	}

	assert.Greater(t, len(meta.GetProjectNames()), 0)
}

func TestGetProjectVersionMeta(t *testing.T) {
	meta := loadMetaWrapper(t)

	projectNames := meta.GetProjectNames()
	projectName := projectNames[0]
	projectVersions := meta.GetProjectVersions(projectName)
	projectVersion := projectVersions[0]

	versionMeta, err := meta.GetVersionMeta(projectName, projectVersion)
	if err != nil {
		t.Errorf("failed to get version meta: %v", err)
	}
	assert.Equal(t, projectName, versionMeta.Name)
	assert.Equal(t, projectVersion, versionMeta.SchemaVersion)
}

func loadMetaWrapper(t *testing.T) *HubMetaWrapper {
	b, err := os.ReadFile("./meta.yaml")
	if err != nil {
		t.Errorf("failed to read file: %v", err)
	}

	hubMeta, err := UnmarshalHubMeta(b)
	if err != nil {
		t.Errorf("failed to unmarshal meta: %v", err)
	}
	wrapper := NewHubMetaWrapper(*hubMeta)
	return wrapper
}
