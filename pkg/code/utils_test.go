package code

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tensorleap/leap-cli/pkg/workspace"
)

func TestBundleCodeIntoTempFile(t *testing.T) {
	// Setup temporary directory and files
	tempDir := t.TempDir()
	file1 := filepath.Join(tempDir, "file1.txt")
	file2 := filepath.Join(tempDir, "file2.txt")
	file3 := filepath.Join(tempDir, "leap_mapping.yaml")
	err := os.WriteFile(file1, []byte("content1"), 0644)
	assert.NoError(t, err)
	err = os.WriteFile(file2, []byte("content2"), 0644)
	assert.NoError(t, err)
	err = os.WriteFile(file3, []byte("mapping content"), 0644)
	assert.NoError(t, err)

	// Setup workspace config
	workspaceConfig := &workspace.WorkspaceConfig{
		IncludePatterns: []string{"*.txt"},
	}

	// Test without externalLeapMappingPath
	closeFunc, tarGzFile, err := BundleCodeIntoTempFile(tempDir, workspaceConfig, "")
	assert.NoError(t, err)
	defer closeFunc()
	defer tarGzFile.Close()

	// Verify tar.gz file is created
	assert.FileExists(t, tarGzFile.Name())

	// Test with externalLeapMappingPath
	externalLeapMappingPath := filepath.Join(tempDir, "external_leap_mapping.yaml")
	err = os.WriteFile(externalLeapMappingPath, []byte("mapping content"), 0644)
	assert.NoError(t, err)

	closeFunc, tarGzFile, err = BundleCodeIntoTempFile(tempDir, workspaceConfig, externalLeapMappingPath)
	assert.NoError(t, err)
	defer closeFunc()
	defer tarGzFile.Close()

	// Verify tar.gz file is created
	assert.FileExists(t, tarGzFile.Name())

	// assert tar.gz file contains the expected files
	filePaths, err := extractTarGzInMemory(tarGzFile)
	assert.NoError(t, err)
	assert.Contains(t, filePaths, "file1.txt")
	assert.Contains(t, filePaths, "file2.txt")
	assert.Contains(t, filePaths, "leap_mapping.yaml")
}

func extractTarGzInMemory(tarGzReader io.Reader) ([]string, error) {
	var filePaths []string

	// Create a gzip reader
	gzipReader, err := gzip.NewReader(tarGzReader)
	if err != nil {
		return nil, err
	}
	defer gzipReader.Close()

	// Create a tar reader
	tarReader := tar.NewReader(gzipReader)

	// Iterate through the tar archive
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break // End of archive
		}
		if err != nil {
			return nil, err
		}

		// Collect file paths
		filePaths = append(filePaths, header.Name)
	}

	return filePaths, nil
}
