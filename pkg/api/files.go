package api

import (
	"fmt"
	"io"
	"mime"
	"net/http"
	"net/url"
	"path"

	"github.com/tensorleap/leap-cli/pkg/log"
)

func DownloadFile(url string, file io.Writer) error {
	log.Println("Downloading...")

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download file. Status code: %d", resp.StatusCode)
	}

	// Create a pipe to copy data concurrently
	reader, writer := io.Pipe()
	go func() {
		defer writer.Close()
		_, err := io.Copy(writer, resp.Body)
		if err != nil {
			writer.CloseWithError(err)
		}
	}()

	_, err = io.Copy(file, reader)
	if err != nil {
		return err
	}

	return nil
}

func UploadFile(url string, file io.Reader, fileSize int64) error {
	log.Infof("Upload target: %s", url)
	s := log.NewSpinner("Uploading...")
	s.Start()
	defer s.Stop()

	req, err := http.NewRequest(http.MethodPut, url, file)
	if err != nil {
		return err
	}
	if fileSize > 0 {
		req.ContentLength = fileSize
	}
	req.Header.Set("Content-Type", DetectContentTypeFromUrl(url))
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err := CheckRes(resp, err); err != nil {
		return fmt.Errorf("failed to upload file: %v", err)
	}
	defer resp.Body.Close()
	return nil
}

const defaultContentType = "application/octet-stream"

func DetectContentTypeFromUrl(u string) string {
	parsedUrl, err := url.Parse(u)
	if err != nil {
		return defaultContentType
	}
	fileExt := path.Ext(parsedUrl.Path)
	mimeType := mime.TypeByExtension(fileExt)
	if mimeType == "" {
		return defaultContentType
	}
	return mimeType
}
