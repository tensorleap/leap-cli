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

const LEAP_SKIP_SSL_VERIFY = "LEAP_SKIP_SSL_VERIFY"

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

	if os.Getenv(LEAP_SKIP_SSL_VERIFY) == "true" {
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

const API_PATH = "/api/v2"
const API_CLOUD_SUBDOMAIN = "://api."

func isCloudUrl(url string) bool {
	return strings.Contains(url, ".tensorleap.ai")
}

func NormalizeAPIUrl(url string) (string, error) {
	url = strings.TrimSuffix(url, "/")
	hasProtocol := strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")
	if !hasProtocol {
		return "", fmt.Errorf("URL must start with http:// or https://")
	}
	isCloud := isCloudUrl(url)
	hasApiSubdomain := strings.Contains(url, API_CLOUD_SUBDOMAIN)
	if isCloud && !hasApiSubdomain {
		url = strings.Replace(url, "://", API_CLOUD_SUBDOMAIN, 1)
	}

	if strings.HasSuffix(url, API_PATH) {
		return url, nil
	}
	return fmt.Sprintf("%s%s", url, API_PATH), nil
}

func ChangeToUIUrl(apiUrl string) string {
	url := strings.Replace(apiUrl, API_PATH, "", 1)
	return strings.Replace(url, API_CLOUD_SUBDOMAIN, "://", 1)
}
