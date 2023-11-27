package api

import (
	"fmt"
	"io"
	"net/http"
)

// HTTPError represents an error with details about an HTTP response.
type HTTPError struct {
	URL        string
	StatusCode int
	Body       string
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("Error: HTTP status code %d\nURL: %s\nResponse body:\n%s", e.StatusCode, e.URL, e.Body)
}

func CheckRes(response *http.Response, err error) error {
	if err != nil {
		return err
	}
	if IsValidStatus(response) {
		return nil
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		body = []byte(fmt.Sprintf("Error reading response body: %s", err))
	}
	return &HTTPError{
		URL:        response.Request.URL.String(),
		StatusCode: response.StatusCode,
		Body:       string(body),
	}
}

func IsValidStatus(response *http.Response) bool {
	return response.StatusCode >= 200 && response.StatusCode < 300
}
