package local

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"os"
)

func CreateTarGzFile(filePaths []string, file *os.File) error {
	fmt.Println("Packing files...")
	gzip := gzip.NewWriter(file)
	defer gzip.Close()
	tarWriter := tar.NewWriter(gzip)
	defer tarWriter.Close()

	for _, filePath := range filePaths {
		fileInfo, err := os.Stat(filePath)
		if err != nil {
			return err
		}
		fmt.Println(filePath)
		header, err := tar.FileInfoHeader(fileInfo, fileInfo.Name())
		if err != nil {
			return err
		}
		header.Name = filePath
		err = tarWriter.WriteHeader(header)
		if err != nil {
			return err
		}
		fileContents, err := os.ReadFile(filePath)
		if err != nil {
			return err
		}
		tarWriter.Write(fileContents)
	}

	return nil
}

func CleanupTempFile(file *os.File) {
	file.Close()
	os.Remove(file.Name())
}
