package api

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// HTTPError represents an error with details about an HTTP response.
type HTTPError struct {
	URL        string
	StatusCode int
	Body       string
}

var ErrAuth = errors.New("user is not currently authenticated, to authenticate - please run 'leap auth login'")

func (e *HTTPError) Error() string {
	if e.StatusCode == 401 {
		return ErrAuth.Error()
	}
	return fmt.Sprintf("Error: HTTP status code %d\nURL: %s\nResponse body:\n%s", e.StatusCode, e.URL, e.Body)
}

type CustomRoundTripper struct {
	originalRoundTripper http.Transport
}

func (c *CustomRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	thereIsNoDomain := strings.HasPrefix(req.URL.String(), "/")
	if thereIsNoDomain {
		return nil, ErrAuth
	}
	res, err := c.originalRoundTripper.RoundTrip(req)
	if err != nil {
		return res, err
	}
	if res.StatusCode == 401 {
		return res, ErrAuth
	}
	return res, nil
}

// NewDefaultClient returns a new http.Client with default settings.
func NewDefaultClient() *http.Client {
	customHttpClient := &http.Client{}

	roundTripper := &CustomRoundTripper{
		originalRoundTripper: http.Transport{},
	}
	customHttpClient.Transport = roundTripper

	if os.Getenv("LEAP_SKIP_SSL_VERIFY") == "true" {
		// skip SSL verification
		roundTripper.originalRoundTripper.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	return customHttpClient
}

func CheckRes(response *http.Response, err error) error {
	if err != nil {
		return err
	}
	if IsValidStatus(response) {
		return nil
	}

	if response.StatusCode == 401 {
		return &HTTPError{
			URL:        response.Request.URL.String(),
			StatusCode: response.StatusCode,
			Body:       ErrAuth.Error(),
		}
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

func init() {
	http.DefaultClient = NewDefaultClient()
}
