package analytics

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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
}

// callMixpanelAlias calls Mixpanel's track API with alias event to link whoami with user_id
func callMixpanelAlias(whoami, userID string) error {
	// Validate input parameters
	if whoami == "" || whoami == "unknown" {
		return fmt.Errorf("invalid whoami value: %s", whoami)
	}
	if userID == "" {
		return fmt.Errorf("invalid userID value: %s", userID)
	}

	// Mixpanel alias is handled through the track API with special properties
	aliasData := map[string]interface{}{
		"token":       MixpanelToken,
		"distinct_id": whoami,
		"$set": map[string]interface{}{
			"$alias": userID,
		},
		"time": time.Now().Unix(),
	}

	aliasJSON, err := json.Marshal(aliasData)
	if err != nil {
		return fmt.Errorf("failed to marshal alias data: %w", err)
	}

	// Mixpanel expects the data to be base64 encoded
	encodedData := []byte(fmt.Sprintf("data=%s", aliasJSON))

	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Send the request to the track endpoint
	resp, err := client.Post(MixpanelEndpoint, "application/x-www-form-urlencoded", bytes.NewBuffer(encodedData))
	if err != nil {
		return fmt.Errorf("failed to send alias to Mixpanel: %w", err)
	}
	defer resp.Body.Close()

	// Read response body for better error debugging
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Mixpanel alias API returned status: %d, body: %s", resp.StatusCode, string(body))
	}

	return nil
}

// saveUserID saves a user ID to the persistent file and sends alias to Mixpanel
func saveUserID(userID string) error {
	// Get current username (whoami)
	currentUsername := getCurrentUsername()

	// Send alias to Mixpanel to link whoami with user_id
	if currentUsername != "unknown" {
		if err := callMixpanelAlias(currentUsername, userID); err != nil {
			// Log the error but don't fail the save operation
			// This ensures the user_id is still saved locally
			fmt.Printf("Warning: Failed to send Mixpanel alias: %v\n", err)
		}
	}

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
	properties["whoami"] = getCurrentUsername()

	// Check if user_id is provided in properties
	if userID, exists := properties["user_id"]; exists {
		userIDStr := fmt.Sprintf("%v", userID)
		savedUserID := getUserID()

		// If we have a different saved user_id, update it
		if savedUserID != "" && savedUserID != userIDStr {
			// Update the saved user_id
			if err := saveUserID(userIDStr); err == nil {
				properties["distinct_id"] = userID
			} else {
				// If we can't save, still use the provided user_id
				properties["distinct_id"] = userID
			}
		} else if savedUserID == "" {
			// No saved user_id, save the new one
			if err := saveUserID(userIDStr); err == nil {
				properties["distinct_id"] = userID
			} else {
				// If we can't save, still use the provided user_id
				properties["distinct_id"] = userID
			}
		} else {
			// Same user_id, use it
			properties["distinct_id"] = userID
		}
	} else {
		// Try to get saved user_id, fallback to current username
		if savedUserID := getUserID(); savedUserID != "" {
			properties["distinct_id"] = savedUserID
		} else {
			properties["distinct_id"] = getCurrentUsername()
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
		return fmt.Errorf("Mixpanel API returned status: %d", resp.StatusCode)
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
