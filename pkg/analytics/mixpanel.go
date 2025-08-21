package analytics

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

const (
	MixpanelEndpoint = "https://api.mixpanel.com/track"
	MixpanelToken   = "f1bf46fb339d8c2930cde8c1acf65491"
)

// deviceIDFile stores the path to the device ID file
var deviceIDFile string

// init initializes the device ID file path
func init() {
	// Get user's home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		// Fallback to current directory if home directory is not accessible
		homeDir = "."
	}
	
	// Create .tensorleap directory if it doesn't exist
	tensorleapDir := filepath.Join(homeDir, ".tensorleap")
	if err := os.MkdirAll(tensorleapDir, 0755); err != nil {
		// Fallback to current directory if we can't create the directory
		tensorleapDir = "."
	}
	
	deviceIDFile = filepath.Join(tensorleapDir, "device_id")
}

// generateNewDeviceID creates a new random device ID
func generateNewDeviceID() string {
	// Generate 16 random bytes
	randomBytes := make([]byte, 16)
	rand.Read(randomBytes)
	return fmt.Sprintf("%x", randomBytes)
}

// getDeviceID retrieves or generates a persistent device ID
func getDeviceID() string {
	// Try to read existing device ID
	if data, err := os.ReadFile(deviceIDFile); err == nil && len(data) > 0 {
		deviceID := string(data)
		// Validate that it's a valid hex string of correct length
		if len(deviceID) == 32 && isValidHex(deviceID) {
			return deviceID
		}
	}
	
	// Generate new device ID
	deviceID := generateNewDeviceID()
	
	// Store the new device ID
	if err := os.WriteFile(deviceIDFile, []byte(deviceID), 0600); err != nil {
		// If we can't write to file, just return the generated ID
		// It won't be persistent but at least it will work for this session
	}
	
	return deviceID
}

// isValidHex checks if a string is a valid hexadecimal string
func isValidHex(s string) bool {
	for _, c := range s {
		if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')) {
			return false
		}
	}
	return true
}

// TrackServerInstall sends a server installation event to Mixpanel
func TrackServerInstall(properties map[string]interface{}) error {
	if properties == nil {
		properties = make(map[string]interface{})
	}

	properties["token"] = MixpanelToken
	properties["time"] = time.Now().Unix()
	properties["os"] = runtime.GOOS
	properties["arch"] = runtime.GOARCH
	properties["timestamp"] = time.Now().Format(time.RFC3339)
	properties["distinct_id"] = getDeviceID()
	properties["device_id"] = getDeviceID()

	if username := os.Getenv("USER"); username != "" {
		properties["os_username"] = username
	}

	event := map[string]interface{}{
		"event":      "server_install",
		"properties": properties,
	}

	eventData, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to marshal event: %w", err)
	}

	// Mixpanel expects the data to be base64 encoded
	encodedData := []byte(fmt.Sprintf("data=%s", eventData))

	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Post(MixpanelEndpoint, "application/x-www-form-urlencoded", bytes.NewBuffer(encodedData))
	if err != nil {
		return fmt.Errorf("failed to send event to Mixpanel: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Mixpanel API returned status: %d", resp.StatusCode)
	}

	return nil
}
