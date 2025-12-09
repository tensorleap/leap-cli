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
	files, err := ExtractTarGzFile(file, outputDir, "")
	assert.NoError(t, err)
	assert.FileExists(t, filepath.Join(outputDir, "project.json"))
	assert.Len(t, files, 13)
}

func TestExtractTarGzFileSpecificFile(t *testing.T) {
	file, _ := os.Open("./test_assets/project.tar.zip")
	defer file.Close()
	outputDir := "./exported"
	defer os.RemoveAll(outputDir)
	files, err := ExtractTarGzFile(file, outputDir, "project.json")
	assert.NoError(t, err)
	assert.FileExists(t, filepath.Join(outputDir, "project.json"))
	assert.Len(t, files, 1)
}

func TestCalculateSHA256(t *testing.T) {
	file, _ := os.Open("./test_assets/project.tar.zip")
	defer file.Close()

	hash1, _ := GetFileChecksum(file)
	_, _ = file.Seek(0, 0)
	hash2, _ := GetFileChecksum(file)

	assert.Equal(t, hash1, hash2)
}

func TestDownloadAndExtractTarFile(t *testing.T) {
	t.Skip("skip test") // fix assets
	url := "https://hub.tensorleap.ai/demo/projects/MNIST/versions/81/1686477167629-project.tar"
	outputDir := "./exported"
	defer os.RemoveAll(outputDir)
	files, err := DownloadAndExtractTarFile(url, outputDir, "")
	assert.NoError(t, err)
	assert.FileExists(t, filepath.Join(outputDir, "project.json"))
	assert.Len(t, files, 42)
}

func TestConvertPathToUnixWithEscaping(t *testing.T) {
	assert.Equal(t, "a/b/c", ConvertPathPatternToUnix("a\\b\\c"))
	assert.Equal(t, "a/b/c", ConvertPathPatternToUnix("a/b/c"))
	assert.Equal(t, "a/b/c", ConvertPathPatternToUnix("a/b\\c"))
	assert.Equal(t, "a/b/c\\[d\\]*", ConvertPathPatternToUnix("a\\b\\c\\[d\\]*"))
}

func TestIsHiddenPath(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected bool
	}{
		{
			name:     "regular file",
			path:     "file.txt",
			expected: false,
		},
		{
			name:     "regular nested file",
			path:     "dir/subdir/file.txt",
			expected: false,
		},
		{
			name:     "hidden file in root",
			path:     ".gitignore",
			expected: true,
		},
		{
			name:     "hidden directory",
			path:     ".git/config",
			expected: true,
		},
		{
			name:     "nested hidden directory",
			path:     ".git/objects/abc/def",
			expected: true,
		},
		{
			name:     "hidden file in regular directory",
			path:     "dir/.hidden",
			expected: true,
		},
		{
			name:     "regular file in hidden directory",
			path:     "dir/.hidden/file.txt",
			expected: true,
		},
		{
			name:     "Windows-style path with hidden folder",
			path:     "dir\\.git\\config",
			expected: true,
		},
		{
			name:     "Windows-style regular path",
			path:     "dir\\subdir\\file.txt",
			expected: false,
		},
		{
			name:     "current directory marker",
			path:     "./file.txt",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isHiddenPath(tt.path)
			assert.Equal(t, tt.expected, result, "Path: %s", tt.path)
		})
	}
}

func TestCreateTarGzFileExcludesHiddenFiles(t *testing.T) {
	// Create a temporary directory with test files
	tempDir := t.TempDir()

	// Create regular files
	regularFile := filepath.Join(tempDir, "regular.txt")
	err := os.WriteFile(regularFile, []byte("regular content"), 0644)
	assert.NoError(t, err)

	nestedDir := filepath.Join(tempDir, "subdir")
	err = os.MkdirAll(nestedDir, 0755)
	assert.NoError(t, err)

	nestedFile := filepath.Join(nestedDir, "nested.txt")
	err = os.WriteFile(nestedFile, []byte("nested content"), 0644)
	assert.NoError(t, err)

	// Create hidden files (should be excluded)
	hiddenFile := filepath.Join(tempDir, ".hidden")
	err = os.WriteFile(hiddenFile, []byte("hidden content"), 0644)
	assert.NoError(t, err)

	gitDir := filepath.Join(tempDir, ".git")
	err = os.MkdirAll(gitDir, 0755)
	assert.NoError(t, err)

	gitConfig := filepath.Join(gitDir, "config")
	err = os.WriteFile(gitConfig, []byte("git config"), 0644)
	assert.NoError(t, err)

	// Create file list including both regular and hidden files
	filePaths := []string{
		"regular.txt",
		"subdir/nested.txt",
		".hidden",
		".git/config",
	}

	// Create tar.gz file
	tarFile, err := os.CreateTemp("", "test-*.tar.gz")
	assert.NoError(t, err)
	defer os.Remove(tarFile.Name())
	defer tarFile.Close()

	err = CreateTarGzFile(tempDir, filePaths, tarFile)
	assert.NoError(t, err)

	// Extract and verify only non-hidden files are included
	_, err = tarFile.Seek(0, 0)
	assert.NoError(t, err)

	extractDir := t.TempDir()
	extractedFiles, err := ExtractTarGzFile(tarFile, extractDir, "")
	assert.NoError(t, err)

	// Should only have 2 non-hidden files
	assert.Len(t, extractedFiles, 2)
	assert.FileExists(t, filepath.Join(extractDir, "regular.txt"))
	assert.FileExists(t, filepath.Join(extractDir, "subdir", "nested.txt"))

	// Hidden files should not exist
	assert.NoFileExists(t, filepath.Join(extractDir, ".hidden"))
	assert.NoFileExists(t, filepath.Join(extractDir, ".git", "config"))
}
