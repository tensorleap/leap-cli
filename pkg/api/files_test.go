package api

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDownloadFile(t *testing.T) {
	url := "https://hub.tensorleap.ai/demo/projects/MNIST/versions/81/1686477167629-project.tar"
	outputFilePath := "TestDownloadFile.tar.zip"
	file, err := os.Create(outputFilePath)
	defer os.Remove(outputFilePath)
	assert.NoError(t, err)
	defer file.Close()
	err = DownloadFile(url, file)
	assert.NoError(t, err)
}
