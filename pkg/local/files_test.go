package local

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractTarGzFile(t *testing.T) {
	file, _ := os.Open("./test_assets/project.tar.zip")
	defer file.Close()
	outputDir := "./exported"
	defer os.RemoveAll(outputDir)
	files, err := ExtractTarGzFile(file, outputDir)
	assert.NoError(t, err)
	assert.FileExists(t, filepath.Join(outputDir, "project.json"))
	assert.Len(t, files, 13)
}

func TestDownloadAndExtractTarFile(t *testing.T) {
	t.Skip("skip test") // fix assets
	url := "https://hub.tensorleap.ai/demo/projects/MNIST/versions/81/1686477167629-project.tar"
	outputDir := "./exported"
	defer os.RemoveAll(outputDir)
	files, err := DownloadAndExtractTarFile(url, outputDir)
	assert.NoError(t, err)
	assert.FileExists(t, filepath.Join(outputDir, "project.json"))
	assert.Len(t, files, 42)
}
