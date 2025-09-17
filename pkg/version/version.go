package version

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/tensorleap/leap-cli/pkg/api"
)

// Meant to be replaced during build with LDFLAG: -X 'github.com/tensorleap/leap-cli/pkg/version.CliVersion=${GIT_TAG}'
var CliVersion = "dev"

const (
	// GitHubRepoURL is the repository URL for the leap-cli
	GitHubRepoURL = "https://github.com/tensorleap/leap-cli"
	// GitHubAPIBaseURL is the base URL for GitHub API
	GitHubAPIBaseURL = "https://api.github.com"
)

// GitHubRelease represents a GitHub release
type GitHubRelease struct {
	TagName     string    `json:"tag_name"`
	Name        string    `json:"name"`
	PublishedAt time.Time `json:"published_at"`
	Prerelease  bool      `json:"prerelease"`
	Draft       bool      `json:"draft"`
}

// GetLatestVersion fetches the latest release version from GitHub
func GetLatestVersion() (string, error) {
	client := api.NewDefaultClient()
	client.Timeout = 10 * time.Second

	// Use the GitHub API to get the latest release
	url := fmt.Sprintf("%s/repos/tensorleap/leap-cli/releases/latest", GitHubAPIBaseURL)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	// Add headers to avoid rate limiting
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("User-Agent", "leap-cli")

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to fetch latest version: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("GitHub API returned status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	var release GitHubRelease
	if err := json.Unmarshal(body, &release); err != nil {
		return "", fmt.Errorf("failed to parse release data: %w", err)
	}

	return release.TagName, nil
}

// GetLatestVersionFromRedirect fetches the latest version using the redirect method (like the install script)
// This method follows redirects to get the latest release URL and extracts the version from it
func GetLatestVersionFromRedirect() (string, error) {
	client := api.NewDefaultClient()
	client.Timeout = 10 * time.Second
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		// Allow up to 10 redirects
		if len(via) >= 10 {
			return fmt.Errorf("too many redirects")
		}
		return nil
	}

	// Use the same URL as the install script
	url := fmt.Sprintf("%s/releases/latest", GitHubRepoURL)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("User-Agent", "leap-cli")

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to fetch latest version: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("GitHub returned status %d", resp.StatusCode)
	}

	// Extract version from the final URL
	finalURL := resp.Request.URL.String()

	// The URL should be something like: https://github.com/tensorleap/leap-cli/releases/tag/v1.2.3
	// We need to extract the version from the end of the URL
	parts := strings.Split(finalURL, "/")
	if len(parts) == 0 {
		return "", fmt.Errorf("invalid redirect URL: %s", finalURL)
	}

	version := parts[len(parts)-1]

	return version, nil
}
