package version

import (
	"testing"
)

func TestGetLatestVersion(t *testing.T) {
	version, err := GetLatestVersion()
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
	version2, err2 := GetLatestVersionFromRedirect()

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
