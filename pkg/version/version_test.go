package version

import (
	"strings"
	"testing"
)

func skipOnRateLimit(t *testing.T, err error) {
	t.Helper()
	if err != nil && strings.Contains(err.Error(), "status 403") {
		t.Skip("Skipping: GitHub API rate limit exceeded (unauthenticated)")
	}
}

func TestGetLatestVersion(t *testing.T) {
	version, err := GetLatestVersion()
	skipOnRateLimit(t, err)
	if err != nil {
		t.Fatalf("GetLatestVersion() failed: %v", err)
	}

	if version == "" {
		t.Error("GetLatestVersion() returned empty version")
	}

	t.Logf("Latest version: %s", version)
}

func TestGetLatestVersionFromRedirect(t *testing.T) {
	version, err := GetLatestVersionFromRedirect()
	skipOnRateLimit(t, err)
	if err != nil {
		t.Fatalf("GetLatestVersionFromRedirect() failed: %v", err)
	}

	if version == "" {
		t.Error("GetLatestVersionFromRedirect() returned empty version")
	}

	t.Logf("Latest version (from redirect): %s", version)
}

func TestGetLatestVersionConsistency(t *testing.T) {
	version1, err1 := GetLatestVersion()
	skipOnRateLimit(t, err1)
	version2, err2 := GetLatestVersionFromRedirect()
	skipOnRateLimit(t, err2)

	if err1 != nil {
		t.Fatalf("GetLatestVersion() failed: %v", err1)
	}

	if err2 != nil {
		t.Fatalf("GetLatestVersionFromRedirect() failed: %v", err2)
	}

	if version1 != version2 {
		t.Errorf("Version mismatch: GetLatestVersion()=%s, GetLatestVersionFromRedirect()=%s", version1, version2)
	}

	t.Logf("Both methods returned consistent version: %s", version1)
}
