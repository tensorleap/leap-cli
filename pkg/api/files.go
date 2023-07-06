package api

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func UploadFile(url string, file *os.File) error {
	fmt.Println("Uploading...")
	reader, writer := io.Pipe()
	go func() {
		defer writer.Close()
		_, err := io.Copy(writer, file)
		writer.CloseWithError(err)
	}()
	req, err := http.NewRequest(http.MethodPut, url, reader)
	if err != nil {
		return err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
