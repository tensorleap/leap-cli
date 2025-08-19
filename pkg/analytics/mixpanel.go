package analytics

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/google/uuid"
	"github.com/spf13/viper"
	"github.com/tensorleap/leap-cli/pkg/version"
)

const (
	// Mixpanel endpoint
	mixpanelEndpoint = "https://api.mixpanel.com/track"
	
	// Event names
	EventCLIInstallStart     = "cli_install_start"
	EventCLIInstallEnd       = "cli_install_end"
	EventCLIInstallFailed    = "cli_install_failed"
	EventCLIReinstall       = "cli_reinstall"
	EventCLIReinstallFailed = "cli_reinstall_failed"
	EventServerInstallStart  = "server_install_start"
	EventServerInstallEnd    = "server_install_end"
	EventServerInstallFailed = "server_install_failed"
	EventServerReinstall    = "server_reinstall"
	EventServerReinstallFailed = "server_reinstall_failed"
	EventServerUpgradeStart  = "server_upgrade_start"
	EventServerUpgradeEnd    = "server_upgrade_end"
	EventServerUpgradeFailed = "server_upgrade_failed"
	EventCLIAuth             = "cli_auth"
	EventCLIAuthFailed       = "cli_auth_failed"
	EventCLICodeInit         = "cli_code_init"
	EventCLICodeInitFailed   = "cli_code_init_failed"
	EventCustomTestRunStart  = "leap_custom_test_run_start"
	EventCustomTestRunEnd    = "leap_custom_test_run_end"
	EventCustomTestRunFailed = "leap_custom_test_run_failed"
	EventCLICodePushStart    = "cli_code_push_start"
	EventCLICodePushFailed   = "cli_code_push_failed"
	EventDatasetParseStart   = "dataset_parse_start"
	EventDatasetParseFailed  = "dataset_parse_failed"
	EventModelsImportStart   = "cli_models_import_start"
	EventModelsImportFailed  = "cli_models_import_failed"
	EventProjectsPush        = "cli_projects_push"
	EventProjectsPushFailed  = "cli_projects_push_failed"
	EventValidateAssetsStart = "validate_assets_start"
	EventValidateAssetsFailed = "validate_assets_failed"
	EventEvaluateStart       = "evaluate_start"
	EventEvaluateFailed      = "evaluate_failed"
	EventPopExplorerStart    = "pop_explorer_start"
	EventPopExplorerFailed   = "pop_explorer_failed"
	EventVisualizerStart     = "visualizer_start"
	EventVisualizerFailed    = "visualizer_failed"
)

// MixpanelClient handles all analytics tracking
type MixpanelClient struct {
	token     string
	machineID string
	cliID     string
	userEmail string
}

// Event represents a Mixpanel event
type Event struct {
	Event      string                 `json:"event"`
	Properties map[string]interface{} `json:"properties"`
}

// NewMixpanelClient creates a new Mixpanel client
func NewMixpanelClient() *MixpanelClient {
	return &MixpanelClient{
		token:     "f1bf46fb339d8c2930cde8c1acf65491", // Your Mixpanel token
		machineID: getOrCreateMachineID(),
		cliID:     getOrCreateCLIID(),
		userEmail: getUserEmail(),
	}
}

// getOrCreateMachineID gets or creates a unique machine identifier
func getOrCreateMachineID() string {
	configKey := "analytics.machine_id"
	
	// Try to get from config first
	if machineID := viper.GetString(configKey); machineID != "" {
		return machineID
	}
	
	// Generate new machine ID
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}
	
	// Create a unique machine ID using hostname and some system info
	machineID := fmt.Sprintf("%s-%s-%s", hostname, runtime.GOOS, runtime.GOARCH)
	
	// Store in config
	viper.Set(configKey, machineID)
	viper.WriteConfig()
	
	return machineID
}

// getOrCreateCLIID gets or creates a unique CLI installation ID
func getOrCreateCLIID() string {
	configKey := "analytics.cli_id"
	
	// Try to get from config first
	if cliID := viper.GetString(configKey); cliID != "" {
		return cliID
	}
	
	// Generate new CLI ID
	cliID := uuid.New().String()
	
	// Store in config
	viper.Set(configKey, cliID)
	viper.WriteConfig()
	
	return cliID
}

// getUserEmail gets the user email from config
func getUserEmail() string {
	return viper.GetString("auth.email")
}

// SetUserEmail updates the user email for analytics
func (m *MixpanelClient) SetUserEmail(email string) {
	m.userEmail = email
	viper.Set("auth.email", email)
	viper.WriteConfig()
}

// Track sends an event to Mixpanel
func (m *MixpanelClient) Track(eventName string, properties map[string]interface{}) {
	// Add common properties
	if properties == nil {
		properties = make(map[string]interface{})
	}
	
	properties["token"] = m.token
	properties["distinct_id"] = m.userEmail
	properties["time"] = time.Now().Unix()
	properties["cli_version"] = version.CliVersion
	properties["machine_id"] = m.machineID
	properties["cli_id"] = m.cliID
	properties["os"] = runtime.GOOS
	properties["arch"] = runtime.GOARCH
	
	event := Event{
		Event:      eventName,
		Properties: properties,
	}
	
	// Encode as base64
	eventData, _ := json.Marshal(event)
	encodedData := base64.StdEncoding.EncodeToString(eventData)
	
	// Send to Mixpanel
	go m.sendToMixpanel(encodedData)
}

// sendToMixpanel sends the event data to Mixpanel
func (m *MixpanelClient) sendToMixpanel(data string) {
	url := fmt.Sprintf("%s?data=%s", mixpanelEndpoint, data)
	
	resp, err := http.Post(url, "text/plain", bytes.NewBufferString(""))
	if err != nil {
		// Silently fail - analytics shouldn't break the CLI
		return
	}
	defer resp.Body.Close()
}

// CLI Install Events
func (m *MixpanelClient) TrackCLIInstallStart() {
	m.Track(EventCLIInstallStart, map[string]interface{}{
		"cli_install_id": m.cliID,
		"cli_version":    version.CliVersion,
		"machine_id":     m.machineID,
	})
}

func (m *MixpanelClient) TrackCLIInstallEnd(serverInstalled bool) {
	m.Track(EventCLIInstallEnd, map[string]interface{}{
		"cli_install_id": m.cliID,
		"cli_version":    version.CliVersion,
		"machine_id":     m.machineID,
		"server_install": serverInstalled,
	})
}

func (m *MixpanelClient) TrackCLIInstallFailed(errorMessage string, errorType string) {
	m.Track(EventCLIInstallFailed, map[string]interface{}{
		"cli_install_id": m.cliID,
		"cli_version":    version.CliVersion,
		"machine_id":     m.machineID,
		"error_message":  errorMessage,
		"error_type":     errorType,
	})
}

func (m *MixpanelClient) TrackCLIReinstall(previousVersion string, reason string) {
	m.Track(EventCLIReinstall, map[string]interface{}{
		"cli_install_id":  m.cliID,
		"cli_version":     version.CliVersion,
		"previous_version": previousVersion,
		"reason":          reason,
		"machine_id":      m.machineID,
	})
}

func (m *MixpanelClient) TrackCLIReinstallFailed(previousVersion string, reason string, errorMessage string, errorType string) {
	m.Track(EventCLIReinstallFailed, map[string]interface{}{
		"cli_install_id":  m.cliID,
		"cli_version":     version.CliVersion,
		"previous_version": previousVersion,
		"reason":          reason,
		"machine_id":      m.machineID,
		"error_message":   errorMessage,
		"error_type":      errorType,
	})
}

// Server Install Events
func (m *MixpanelClient) TrackServerInstallStart(installType string, tlVersion string) {
	installID := uuid.New().String()
	
	m.Track(EventServerInstallStart, map[string]interface{}{
		"install_type": installType, // "upgrade", "install", "reinstall"
		"install_id":  installID,
		"tl_version":  tlVersion,
		"machine_id":  m.machineID,
	})
}

func (m *MixpanelClient) TrackServerInstallEnd(installID string, timeToInstall time.Duration, availableStorage int64, mountCount int, gpuCount int, gpuTypes []string, downloadSizeGB float64, wasSuccessful bool, machineSetup map[string]interface{}) {
	m.Track(EventServerInstallEnd, map[string]interface{}{
		"install_id":         installID,
		"cli_install_id":     m.cliID,
		"time_to_install":    timeToInstall.Seconds(),
		"available_storage":  availableStorage,
		"mount_count":        mountCount,
		"gpu_count":          gpuCount,
		"gpu_types":          gpuTypes,
		"download_size_gb":   downloadSizeGB,
		"was_successful":     wasSuccessful,
		"machine_setup":      machineSetup,
		"machine_id":         m.machineID,
	})
}

func (m *MixpanelClient) TrackServerReinstall(installID string, previousVersion string, reason string, tlVersion string) {
	m.Track(EventServerReinstall, map[string]interface{}{
		"install_id":      installID,
		"cli_install_id":  m.cliID,
		"previous_version": previousVersion,
		"reason":          reason,
		"tl_version":      tlVersion,
		"machine_id":      m.machineID,
	})
}

func (m *MixpanelClient) TrackServerReinstallFailed(installID string, previousVersion string, reason string, tlVersion string, errorMessage string, errorType string) {
	m.Track(EventServerReinstallFailed, map[string]interface{}{
		"install_id":      installID,
		"cli_install_id":  m.cliID,
		"previous_version": previousVersion,
		"reason":          reason,
		"tl_version":      tlVersion,
		"machine_id":      m.machineID,
		"error_message":   errorMessage,
		"error_type":      errorType,
	})
}

// Server Upgrade Events
func (m *MixpanelClient) TrackServerUpgradeStart(upgradeID string, fromVersion string, toVersion string) {
	m.Track(EventServerUpgradeStart, map[string]interface{}{
		"upgrade_id":    upgradeID,
		"cli_install_id": m.cliID,
		"from_version":  fromVersion,
		"to_version":    toVersion,
		"machine_id":    m.machineID,
	})
}

func (m *MixpanelClient) TrackServerUpgradeEnd(upgradeID string, timeToUpgrade time.Duration, wasSuccessful bool) {
	m.Track(EventServerUpgradeEnd, map[string]interface{}{
		"upgrade_id":      upgradeID,
		"cli_install_id":  m.cliID,
		"time_to_upgrade": timeToUpgrade.Seconds(),
		"was_successful":  wasSuccessful,
		"machine_id":      m.machineID,
	})
}

func (m *MixpanelClient) TrackServerUpgradeFailed(upgradeID string, fromVersion string, toVersion string, errorMessage string, errorType string) {
	m.Track(EventServerUpgradeFailed, map[string]interface{}{
		"upgrade_id":     upgradeID,
		"cli_install_id": m.cliID,
		"from_version":   fromVersion,
		"to_version":     toVersion,
		"machine_id":     m.machineID,
		"error_message":  errorMessage,
		"error_type":     errorType,
	})
}

// CLI Auth Events
func (m *MixpanelClient) TrackCLIAuth(authType string, serverType string, serverURL string) {
	m.Track(EventCLIAuth, map[string]interface{}{
		"cli_install_id": m.cliID,
		"auth_type":      authType, // "leap_auth_select", "copy_paste"
		"server_type":    serverType, // "local", "saas", "azure", "sagemaker"
		"server_url":     serverURL,
	})
}

func (m *MixpanelClient) TrackCLIAuthFailed(authType string, serverType string, serverURL string, errorMessage string, errorType string) {
	m.Track(EventCLIAuthFailed, map[string]interface{}{
		"cli_install_id": m.cliID,
		"auth_type":      authType,
		"server_type":    serverType,
		"server_url":     serverURL,
		"error_message":  errorMessage,
		"error_type":     errorType,
	})
}

// CLI Code Events
func (m *MixpanelClient) TrackCLICodeInit(isNew bool, codeID string) {
	m.Track(EventCLICodeInit, map[string]interface{}{
		"is_new":  isNew,
		"code_id": codeID,
	})
}

func (m *MixpanelClient) TrackCLICodeInitFailed(isNew bool, errorMessage string, errorType string) {
	m.Track(EventCLICodeInitFailed, map[string]interface{}{
		"is_new":        isNew,
		"error_message": errorMessage,
		"error_type":    errorType,
	})
}

// Custom Test Events
func (m *MixpanelClient) TrackCustomTestRunStart(testID string) {
	m.Track(EventCustomTestRunStart, map[string]interface{}{
		"test_id": testID,
	})
}

func (m *MixpanelClient) TrackCustomTestRunEnd(testID string, completedIntegrations int, isSuccessful bool, errorIn string, samplesRun int) {
	m.Track(EventCustomTestRunEnd, map[string]interface{}{
		"test_id":                 testID,
		"completed_integrations":  completedIntegrations,
		"is_successful":           isSuccessful,
		"error_in":                errorIn, // "input", "gt", "preprocess", "visualizer"
		"samples_run":             samplesRun,
	})
}

func (m *MixpanelClient) TrackCustomTestRunFailed(testID string, errorMessage string, errorType string, samplesRun int) {
	m.Track(EventCustomTestRunFailed, map[string]interface{}{
		"test_id":        testID,
		"error_message":  errorMessage,
		"error_type":     errorType,
		"samples_run":    samplesRun,
	})
}

// Code Push Events
func (m *MixpanelClient) TrackCLICodePushStart(projectID string, codeID string, codeVersionID string) {
	m.Track(EventCLICodePushStart, map[string]interface{}{
		"cli_install_id": m.cliID,
		"project_id":     projectID,
		"code_id":        codeID,
		"code_version_id": codeVersionID,
	})
}

func (m *MixpanelClient) TrackCLICodePushFailed(projectID string, codeID string, codeVersionID string, errorMessage string, errorType string) {
	m.Track(EventCLICodePushFailed, map[string]interface{}{
		"cli_install_id": m.cliID,
		"project_id":     projectID,
		"code_id":        codeID,
		"code_version_id": codeVersionID,
		"error_message":  errorMessage,
		"error_type":     errorType,
	})
}

// Dataset Parse Events
func (m *MixpanelClient) TrackDatasetParseStart(codeVersionID string, codeID string, engineVersion string, dependencyOnOff bool) {
	m.Track(EventDatasetParseStart, map[string]interface{}{
		"code_version_id": codeVersionID,
		"code_id":         codeID,
		"engine_version":  engineVersion,
		"dependency_on_off": dependencyOnOff,
	})
}

func (m *MixpanelClient) TrackDatasetParseFailed(codeVersionID string, codeID string, engineVersion string, errorMessage string, errorType string) {
	m.Track(EventDatasetParseFailed, map[string]interface{}{
		"code_version_id": codeVersionID,
		"code_id":         codeID,
		"engine_version":  engineVersion,
		"error_message":   errorMessage,
		"error_type":      errorType,
	})
}

// Models Import Events
func (m *MixpanelClient) TrackModelsImportStart(modelName string, modelType string, transformInput bool, engineVersion string) {
	m.Track(EventModelsImportStart, map[string]interface{}{
		"model_name":      modelName,
		"model_type":      modelType, // "onnx", "h5"
		"transform_input": transformInput,
		"engine_version":  engineVersion,
	})
}

func (m *MixpanelClient) TrackModelsImportFailed(modelName string, modelType string, transformInput bool, engineVersion string, errorMessage string, errorType string) {
	m.Track(EventModelsImportFailed, map[string]interface{}{
		"model_name":      modelName,
		"model_type":      modelType,
		"transform_input": transformInput,
		"engine_version":  engineVersion,
		"error_message":   errorMessage,
		"error_type":      errorType,
	})
}

// Projects Push Events
func (m *MixpanelClient) TrackProjectsPush(projectID string, modelName string, modelType string, transformInput bool, engineVersion string) {
	m.Track(EventProjectsPush, map[string]interface{}{
		"project_id":     projectID,
		"model_name":     modelName,
		"model_type":     modelType,
		"transform_input": transformInput,
		"engine_version":  engineVersion,
	})
}

func (m *MixpanelClient) TrackProjectsPushFailed(projectID string, modelName string, modelType string, transformInput bool, engineVersion string, errorMessage string, errorType string) {
	m.Track(EventProjectsPushFailed, map[string]interface{}{
		"project_id":     projectID,
		"model_name":     modelName,
		"model_type":     modelType,
		"transform_input": transformInput,
		"engine_version":  engineVersion,
		"error_message":  errorMessage,
		"error_type":     errorType,
	})
}

// Validate Assets Events
func (m *MixpanelClient) TrackValidateAssetsStart() {
	m.Track(EventValidateAssetsStart, map[string]interface{}{})
}

func (m *MixpanelClient) TrackValidateAssetsFailed(errorMessage string, errorType string) {
	m.Track(EventValidateAssetsFailed, map[string]interface{}{
		"error_message": errorMessage,
		"error_type":    errorType,
	})
}

// Evaluate Events
func (m *MixpanelClient) TrackEvaluateStart(batch bool, skipMetrics bool, runName string, codeID string, projectID string) {
	m.Track(EventEvaluateStart, map[string]interface{}{
		"batch":        batch,
		"skip_metrics": skipMetrics,
		"run_name":     runName,
		"code_id":      codeID,
		"project_id":   projectID,
	})
}

func (m *MixpanelClient) TrackEvaluateFailed(batch bool, skipMetrics bool, runName string, codeID string, projectID string, errorMessage string, errorType string) {
	m.Track(EventEvaluateFailed, map[string]interface{}{
		"batch":         batch,
		"skip_metrics":  skipMetrics,
		"run_name":      runName,
		"code_id":       codeID,
		"project_id":    projectID,
		"error_message": errorMessage,
		"error_type":    errorType,
	})
}

// Pop Explorer Events
func (m *MixpanelClient) TrackPopExplorerStart(runName string, codeID string, projectID string, filters map[string]interface{}) {
	m.Track(EventPopExplorerStart, map[string]interface{}{
		"run_name":   runName,
		"code_id":    codeID,
		"project_id": projectID,
		"filters":    filters,
	})
}

func (m *MixpanelClient) TrackPopExplorerFailed(runName string, codeID string, projectID string, errorMessage string, errorType string) {
	m.Track(EventPopExplorerFailed, map[string]interface{}{
		"run_name":      runName,
		"code_id":       codeID,
		"project_id":    projectID,
		"error_message": errorMessage,
		"error_type":    errorType,
	})
}

// Visualizer Events
func (m *MixpanelClient) TrackVisualizerStart(runName string, codeID string, projectID string, visID string, samplesCount int, visualizersCount int, visTypeCounts map[string]int, isAuto bool) {
	m.Track(EventVisualizerStart, map[string]interface{}{
		"run_name":           runName,
		"code_id":            codeID,
		"project_id":         projectID,
		"vis_id":             visID,
		"samples_count":      samplesCount,
		"visualizers_count":  visualizersCount,
		"vis_type_counts":    visTypeCounts,
		"is_auto":            isAuto,
	})
}

func (m *MixpanelClient) TrackVisualizerFailed(runName string, codeID string, projectID string, visID string, errorMessage string, errorType string) {
	m.Track(EventVisualizerFailed, map[string]interface{}{
		"run_name":      runName,
		"code_id":       codeID,
		"project_id":    projectID,
		"vis_id":        visID,
		"error_message": errorMessage,
		"error_type":    errorType,
	})
}
