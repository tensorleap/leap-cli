package code

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tensorleap/leap-cli/pkg/workspace"
)

func TestGetCodeFiles(t *testing.T) {
	// Setup temporary directory and files
	tempDir := t.TempDir()
	nestedDir := filepath.Join(tempDir, "nested")
	err := os.Mkdir(nestedDir, 0755)
	assert.NoError(t, err)

	// Create test files
	files := []string{
		"file1.txt",
		"file2.log",
		"pyproject.toml",
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
				"nested/file3.txt",
				"pyproject.toml",
				"nested/file4.log",
			},
		},
		{
			name:            "Include only *.txt",
			includePatterns: []string{"**/*.txt", "pyproject.toml"},
			excludePatterns: nil,
			expectedFiles: []string{
				"file1.txt",
				"nested/file3.txt",
				"pyproject.toml",
			},
		},
		{
			name:            "Exclude *.log",
			includePatterns: nil,
			excludePatterns: []string{"**/*.log"},
			expectedFiles: []string{
				"file1.txt",
				"pyproject.toml",
				"nested/file3.txt",
			},
		},
		{
			name:            "Include *.txt and exclude nested/*",
			includePatterns: []string{"**/*.txt", "pyproject.toml"},
			excludePatterns: []string{"nested/*"},
			expectedFiles: []string{
				"file1.txt",
				"pyproject.toml",
			},
		},
		{
			name:            "Include all and exclude *.log",
			includePatterns: []string{"**"},
			excludePatterns: []string{"**/*.log"},
			expectedFiles: []string{
				"file1.txt",
				"nested/file3.txt",
				"pyproject.toml",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			workspaceConfig := &workspace.WorkspaceConfig{
				IncludePatterns: test.includePatterns,
				ExcludePatterns: test.excludePatterns,
			}

			files, err := getCodeFiles(tempDir, workspaceConfig)
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
	file3 := filepath.Join(tempDir, "pyproject.toml")
	err := os.WriteFile(file1, []byte("content1"), 0644)
	assert.NoError(t, err)
	err = os.WriteFile(file2, []byte("content2"), 0644)
	assert.NoError(t, err)
	err = os.WriteFile(file3, []byte("content3"), 0644)
	assert.NoError(t, err)
	// Setup workspace config
	workspaceConfig := &workspace.WorkspaceConfig{
		IncludePatterns: []string{"*.txt", "pyproject.toml"},
	}

	// Test without externalLeapMappingPath
	closeFunc, tarGzFile, err := BundleCodeIntoTempFile(tempDir, workspaceConfig)
	assert.NoError(t, err)
	defer closeFunc()
	defer tarGzFile.Close()

	// Verify tar.gz file is created
	assert.FileExists(t, tarGzFile.Name())
}

func TestGetCodeFilesExcludesHiddenFiles(t *testing.T) {
	// Setup temporary directory and files
	tempDir := t.TempDir()
	
	// Create regular files
	err := os.WriteFile(filepath.Join(tempDir, "regular.txt"), []byte("regular"), 0644)
	assert.NoError(t, err)
	
	nestedDir := filepath.Join(tempDir, "subdir")
	err = os.Mkdir(nestedDir, 0755)
	assert.NoError(t, err)
	
	err = os.WriteFile(filepath.Join(nestedDir, "nested.txt"), []byte("nested"), 0644)
	assert.NoError(t, err)
	
	// Create hidden files (should be excluded)
	err = os.WriteFile(filepath.Join(tempDir, ".hidden"), []byte("hidden"), 0644)
	assert.NoError(t, err)
	
	gitDir := filepath.Join(tempDir, ".git")
	err = os.Mkdir(gitDir, 0755)
	assert.NoError(t, err)
	
	gitObjectsDir := filepath.Join(gitDir, "objects")
	err = os.Mkdir(gitObjectsDir, 0755)
	assert.NoError(t, err)
	
	err = os.WriteFile(filepath.Join(gitDir, "config"), []byte("git config"), 0644)
	assert.NoError(t, err)
	
	err = os.WriteFile(filepath.Join(gitObjectsDir, "abc"), []byte("object"), 0644)
	assert.NoError(t, err)
	
	// Test with no patterns (should include all non-hidden files)
	workspaceConfig := &workspace.WorkspaceConfig{
		IncludePatterns: nil,
		ExcludePatterns: nil,
	}
	
	files, err := getCodeFiles(tempDir, workspaceConfig)
	assert.NoError(t, err)
	
	// Normalize file paths for comparison
	normalizedFiles := make([]string, len(files))
	for i, file := range files {
		normalizedFiles[i] = filepath.ToSlash(file)
	}
	
	// Should only contain non-hidden files
	expectedFiles := []string{
		"regular.txt",
		"subdir/nested.txt",
	}
	
	assert.ElementsMatch(t, expectedFiles, normalizedFiles, "Hidden files should be excluded")
	
	// Verify hidden files are not in the list
	for _, file := range normalizedFiles {
		assert.NotContains(t, file, ".git", "Should not contain .git files")
		assert.NotContains(t, file, ".hidden", "Should not contain .hidden files")
	}
}
