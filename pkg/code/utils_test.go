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

func TestGetDatasetFiles(t *testing.T) {
	// Setup temporary directory and files
	tempDir := t.TempDir()
	nestedDir := filepath.Join(tempDir, "nested")
	err := os.Mkdir(nestedDir, 0755)
	assert.NoError(t, err)

	// Create test files
	files := []string{
		"file1.txt",
		"file2.log",
		"leap_mapping.yaml",
		"nested/file3.txt",
		"nested/file4.log",
	}
	for _, file := range files {
		path := filepath.Join(tempDir, file)
		err := os.WriteFile(path, []byte("content"), 0644)
		assert.NoError(t, err)
	}

	// Test cases for getDatasetFiles
	tests := []struct {
		name            string
		includePatterns []string
		excludePatterns []string
		expectedFiles   []string
	}{
		{
			name:            "No include or exclude patterns",
			includePatterns: nil,
			excludePatterns: nil,
			expectedFiles: []string{
				"file1.txt",
				"file2.log",
				"leap_mapping.yaml",
				"nested/file3.txt",
				"nested/file4.log",
			},
		},
		{
			name:            "Include only *.txt",
			includePatterns: []string{"**/*.txt"},
			excludePatterns: nil,
			expectedFiles: []string{
				"file1.txt",
				"nested/file3.txt",
				"leap_mapping.yaml",
			},
		},
		{
			name:            "Exclude *.log",
			includePatterns: nil,
			excludePatterns: []string{"**/*.log"},
			expectedFiles: []string{
				"file1.txt",
				"leap_mapping.yaml",
				"nested/file3.txt",
			},
		},
		{
			name:            "Include *.txt and exclude nested/*",
			includePatterns: []string{"**/*.txt"},
			excludePatterns: []string{"nested/*"},
			expectedFiles: []string{
				"file1.txt",
				"leap_mapping.yaml",
			},
		},
		{
			name:            "Include all and exclude *.log",
			includePatterns: []string{"**"},
			excludePatterns: []string{"**/*.log"},
			expectedFiles: []string{
				"file1.txt",
				"leap_mapping.yaml",
				"nested/file3.txt",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			workspaceConfig := &workspace.WorkspaceConfig{
				IncludePatterns: test.includePatterns,
				ExcludePatterns: test.excludePatterns,
			}

			files, err := getDatasetFiles(tempDir, workspaceConfig)
			assert.NoError(t, err)

			// Normalize file paths for comparison
			normalizedFiles := make([]string, len(files))
			for i, file := range files {
				normalizedFiles[i] = filepath.ToSlash(file)
			}

			assert.ElementsMatch(t, test.expectedFiles, normalizedFiles)
		})
	}
}

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
	closeFunc, tarGzFile, err := BundleCodeIntoTempFile(tempDir, workspaceConfig)
	assert.NoError(t, err)
	defer closeFunc()
	defer tarGzFile.Close()

	// Verify tar.gz file is created
	assert.FileExists(t, tarGzFile.Name())
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
