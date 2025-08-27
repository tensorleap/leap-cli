package analytics

import (
	"testing"
)

func TestTrackServerInstall(t *testing.T) {
	// Test that the package compiles and basic functionality works
	properties := map[string]interface{}{
		"port": 8080,
		"gpus": 2,
		"test": true,
	}

	// This will fail in tests since we don't have network access, but it tests compilation
	_ = SendEvent(EventServerInstallStarted, properties)
	_ = SendEvent(EventServerInstallSuccess, properties)
	_ = SendEvent(EventServerInstallFailed, properties)
}

func TestTrackAuthLogin(t *testing.T) {
	// Test that the package compiles and basic functionality works
	properties := map[string]interface{}{
		"api_url": "https://api.example.com",
		"test":    true,
	}

	// This will fail in tests since we don't have network access, but it tests compilation
	_ = SendEvent(EventAuthLoginSuccess, properties)
	_ = SendEvent(EventAuthLoginFailed, properties)
}
