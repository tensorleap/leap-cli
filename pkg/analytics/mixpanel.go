package analytics

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
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
	MixpanelToken    = "f1bf46fb339d8c2930cde8c1acf65491"
)

// EventType represents the type of event to track
type EventType string

const (
	EventCliInstallStarted      EventType = "cli_install_started"
	EventCliInstallSuccess      EventType = "cli_install_success"
	EventCliInstallFailed       EventType = "cli_install_failed"
	EventCliCodeInitStarted     EventType = "cli_code_init_started"
	EventCliCodeInitSuccess     EventType = "cli_code_init_success"
	EventCliCodeInitFailed      EventType = "cli_code_init_failed"
	EventCliCodePushStarted     EventType = "cli_code_push_started"
	EventCliCodePushSuccess     EventType = "cli_code_push_success"
	EventCliCodePushFailed      EventType = "cli_code_push_failed"
	EventCliDatasetParseStarted EventType = "cli_dataset_parse_started"
	EventCliDatasetParseSuccess EventType = "cli_dataset_parse_success"
	EventCliDatasetParseFailed  EventType = "cli_dataset_parse_failed"
	EventCliModelsImportStarted EventType = "cli_models_import_started"
	EventCliModelsImportSuccess EventType = "cli_models_import_success"
	EventCliModelsImportFailed  EventType = "cli_models_import_failed"
	EventCliProjectsPushStarted EventType = "cli_projects_push_started"
	EventCliProjectsPushSuccess EventType = "cli_projects_push_success"
	EventCliProjectsPushFailed  EventType = "cli_projects_push_failed"
	EventAuthLoginSuccess       EventType = "auth_login_success"
	EventAuthLoginFailed        EventType = "auth_login_failed"
	EventAuthLogoutSuccess      EventType = "auth_logout_success"
	EventAuthLogoutFailed       EventType = "auth_logout_failed"
	EventServerInstallStarted   EventType = "server_install_started"
	EventServerInstallSuccess   EventType = "server_install_success"
	EventServerInstallFailed    EventType = "server_install_failed"
)

// userIDFile stores the path to the user ID file
var userIDFile string

// deviceIDFile stores the path to the device ID file
var deviceIDFile string

// init initializes the user ID file path
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

	userIDFile = filepath.Join(tensorleapDir, "user_id")
	deviceIDFile = filepath.Join(tensorleapDir, "device_id")
}

func saveUserID(userID string) error {
	// Save the user ID to the persistent file
	return os.WriteFile(userIDFile, []byte(userID), 0600)
}

// getUserID retrieves the persistent user ID if it exists
func getUserID() string {
	// Try to read existing user ID
	if data, err := os.ReadFile(userIDFile); err == nil && len(data) > 0 {
		userID := string(data)
		// Validate that it's not empty
		if len(userID) > 0 {
			return userID
		}
	}

	// No user ID found, return empty string
	return ""
}

// DeleteUserID removes the persistent user ID file
func DeleteUserID() error {
	return os.Remove(userIDFile)
}

// generateDeviceID generates a new random device ID
func generateDeviceID() (string, error) {
	// Generate 16 random bytes
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	// Convert to hex string
	return hex.EncodeToString(bytes), nil
}

// saveDeviceID saves the device ID to the persistent file
func saveDeviceID(deviceID string) error {
	// Save the device ID to the persistent file
	return os.WriteFile(deviceIDFile, []byte(deviceID), 0600)
}

// getDeviceID retrieves the persistent device ID, generating one if it doesn't exist
func getDeviceID() string {
	// Try to read existing device ID
	if data, err := os.ReadFile(deviceIDFile); err == nil && len(data) > 0 {
		deviceID := string(data)
		// Validate that it's not empty
		if len(deviceID) > 0 {
			return deviceID
		}
	}

	// No device ID found, generate a new one
	deviceID, err := generateDeviceID()
	if err != nil {
		// Fallback to a timestamp-based ID if generation fails
		deviceID = fmt.Sprintf("device_%d", time.Now().UnixNano())
	}

	// Save the new device ID
	if err := saveDeviceID(deviceID); err != nil {
		// If we can't save, still return the generated ID
		return deviceID
	}

	return deviceID
}

// getCurrentUsername gets the current username from environment variables or user info
func getCurrentUsername() string {
	// Try to get from environment variables first (cross-platform)
	if username := os.Getenv("USER"); username != "" {
		return username
	}
	if username := os.Getenv("USERNAME"); username != "" {
		return username
	}

	// Fallback to unknown
	return "unknown"
}

// SendEvent sends an event to Mixpanel with the given event type and properties
func SendEvent(eventType EventType, properties map[string]interface{}) error {
	if properties == nil {
		properties = make(map[string]interface{})
	}

	properties["token"] = MixpanelToken
	properties["time"] = time.Now().Unix()
	properties["os"] = runtime.GOOS
	properties["arch"] = runtime.GOARCH
	properties["$device_id"] = getDeviceID()
	properties["whoami"] = getCurrentUsername()

	// Check if user_id is provided in properties
	if userID, exists := properties["user_id"]; exists {
		properties["$user_id"] = userID
		properties["user_id"] = userID
		userIDStr := fmt.Sprintf("%v", userID)
		savedUserID := getUserID()

		// If we have a different saved user_id, update it
		if savedUserID != "" && savedUserID != userIDStr {
			// Update the saved user_id
			if err := saveUserID(userIDStr); err == nil {
				properties["distinct_id"] = userID
				properties["$user_id"] = userID
			} else {
				// If we can't save, still use the provided user_id
				properties["distinct_id"] = userID
				properties["$user_id"] = userID
			}
		} else if savedUserID == "" {
			// No saved user_id, save the new one
			if err := saveUserID(userIDStr); err == nil {
				properties["distinct_id"] = userID
				properties["$user_id"] = userID
			} else {
				// If we can't save, still use the provided user_id
				properties["distinct_id"] = userID
				properties["$user_id"] = userID
			}
		} else {
			// Same user_id, use it
			properties["distinct_id"] = userID
			properties["$user_id"] = userID
		}
	} else {
		// Try to get saved user_id, fallback to device_id
		if savedUserID := getUserID(); savedUserID != "" {
			properties["$user_id"] = savedUserID
			properties["user_id"] = savedUserID
			properties["distinct_id"] = savedUserID
			properties["$user_id"] = savedUserID
		} else {
			// No userId found, use device_id as distinct_id
			deviceID := getDeviceID()
			properties["distinct_id"] = deviceID
		}
	}

	if username := os.Getenv("USER"); username != "" {
		properties["os_username"] = username
	}

	event := map[string]interface{}{
		"event":      string(eventType),
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
		return fmt.Errorf("mixpanel API returned status: %d", resp.StatusCode)
	}

	return nil
}

// SendEventWithUserData sends an event to Mixpanel and enriches it with user data from a callback
func SendEventWithUserData(eventType EventType, properties map[string]interface{}, userDataFetcher func() map[string]interface{}) error {
	if properties == nil {
		properties = make(map[string]interface{})
	}

	// Try to get user data if fetcher is provided
	if userDataFetcher != nil {
		if userData := userDataFetcher(); userData != nil {
			// Merge user data into properties
			for k, v := range userData {
				properties[k] = v
			}
		}
	}

	return SendEvent(eventType, properties)
}
