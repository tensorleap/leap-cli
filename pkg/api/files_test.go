package api

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDownloadFile(t *testing.T) {
	t.Skip("skip test") // fix assets
	url := "https://hub.tensorleap.ai/demo/projects/MNIST/versions/81/1686477167629-project.tar"
	outputFilePath := "TestDownloadFile.tar.zip"
	file, err := os.Create(outputFilePath)
	defer func() { _ = os.Remove(outputFilePath) }()
	assert.NoError(t, err)
	defer func() { _ = file.Close() }()
	err = DownloadFile(url, file)
	assert.NoError(t, err)
}
