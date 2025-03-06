package local

import (
	"archive/tar"
	"compress/gzip"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/tensorleap/leap-cli/pkg/api"
)

func DownloadAndExtractTarFile(url string, outputDir string, specificFileName string) ([]string, error) {
	reader, writer := io.Pipe()

	downloadErrCh := make(chan error, 1)
	go func() {
		defer writer.Close()
		downloadErrCh <- api.DownloadFile(url, writer)
	}()

	files, extractionErr := ExtractTarGzFile(reader, outputDir, specificFileName)

	downloadErr := <-downloadErrCh

	if extractionErr != nil {
		return nil, fmt.Errorf("failed to extract tar.gz file: %w", extractionErr)
	}
	if downloadErr != nil {
		return nil, fmt.Errorf("failed to download file: %w", downloadErr)
	}

	return files, nil
}

func ExtractTarGzFile(file io.Reader, outputDir string, specificFileName string) ([]string, error) {
	fmt.Println("Extract files...")
	gzipReader, err := gzip.NewReader(file)
	files := []string{}

	if err != nil {
		return files, err
	}
	defer gzipReader.Close()

	tarReader := tar.NewReader(gzipReader)

	// Determine if we are targeting a specific file
	var targetFile string
	if len(specificFileName) > 0 {
		targetFile = filepath.Clean(specificFileName)
	}

	for {
		// Read the next header
		header, err := tarReader.Next()
		if err == io.EOF {
			// End of archive
			break
		}
		if err != nil {
			return files, fmt.Errorf("error reading tar header: %w", err)
		}

		// Clean and normalize the file name
		fileName := filepath.Clean(header.Name)

		// If specific file is provided, skip others
		if targetFile != "" && fileName != targetFile {
			continue
		}

		// Extract the file or directory
		extractedFilePath, err := readNextTarFile(tarReader, header, outputDir)
		if err != nil {
			return files, fmt.Errorf("failed to extract %s: %w", fileName, err)
		}

		files = append(files, ConvertPathToUnix(extractedFilePath))
	}

	return files, nil
}

func readNextTarFile(tarReader *tar.Reader, header *tar.Header, outputDir string) (fileName string, err error) {
	fileName = filepath.Clean(header.Name)
	extractedFilePath := filepath.Join(outputDir, fileName)

	if header.FileInfo().IsDir() {
		// Create the directory
		err = os.MkdirAll(extractedFilePath, header.FileInfo().Mode())
		return
	}

	// Ensure the directory for the file exists
	err = os.MkdirAll(filepath.Dir(extractedFilePath), 0755)
	if err != nil {
		return
	}

	// Create the file and copy the content
	outFile, err := os.Create(extractedFilePath)
	if err != nil {
		return
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, tarReader)
	if err != nil {
		return
	}

	return
}

func addFileToTar(tarWriter *tar.Writer, filePath string, fileName string) error {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return err
	}

	if fileName == "" {
		fileName = fileInfo.Name()
	}

	fmt.Println("Adding:", filePath)
	header, err := tar.FileInfoHeader(fileInfo, fileName)
	if err != nil {
		return err
	}

	header.Name = ConvertPathToUnix(fileName)

	if err := tarWriter.WriteHeader(header); err != nil {
		return err
	}

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(tarWriter, file)
	return err
}

func CreateTarGzFile(filesDir string, filePaths []string, file io.Writer, externalFilePath string, externalFileName string) error {
	fmt.Println("Packing files...")

	gzipWriter := gzip.NewWriter(file)
	defer gzipWriter.Close()

	tarWriter := tar.NewWriter(gzipWriter)
	defer tarWriter.Close()

	for _, relativePath := range filePaths {
		fullFilePath := filepath.Join(filesDir, relativePath)
		if err := addFileToTar(tarWriter, fullFilePath, relativePath); err != nil {
			return err
		}
	}

	if externalFilePath != "" {
		if err := addFileToTar(tarWriter, externalFilePath, externalFileName); err != nil {
			return err
		}
	}

	return nil
}

func ConvertPathToUnix(windowsPath string) string {
	return strings.ReplaceAll(windowsPath, "\\", "/")
}

// Convert only the path separators (backslashes) to forward slashes while preserving escape sequences.
func ConvertPathPatternToUnix(pattern string) string {
	var result strings.Builder

	for i := 0; i < len(pattern); i++ {
		if pattern[i] == '\\' {
			// Check if the next character exists and is a special character (escape sequence)
			if i+1 < len(pattern) && strings.ContainsAny(string(pattern[i+1]), `[]{}()*+?.,\\^$|#`) {
				result.WriteByte(pattern[i])
				result.WriteByte(pattern[i+1])
				i++
			} else {
				result.WriteByte('/')
			}
		} else {
			result.WriteByte(pattern[i])
		}
	}

	return result.String()
}

func CleanupTempFile(file *os.File) {
	file.Close()
	os.Remove(file.Name())
}

func GetFileChecksum(file io.Reader) (string, error) {
	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}
