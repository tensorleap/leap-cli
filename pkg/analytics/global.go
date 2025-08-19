package analytics

import (
	"sync"
	"time"
)

var (
	globalClient *MixpanelClient
	once         sync.Once
)

// GetClient returns the global Mixpanel client instance
func GetClient() *MixpanelClient {
	once.Do(func() {
		globalClient = NewMixpanelClient()
	})
	return globalClient
}

// TrackCLIInstall tracks CLI installation events
func TrackCLIInstall() {
	client := GetClient()
	client.TrackCLIInstallStart()
	
	// Track CLI install end when the program exits
	// This is a simple approach - in a real implementation you might want
	// to use a more sophisticated shutdown mechanism
	go func() {
		time.Sleep(1 * time.Second) // Small delay to ensure start event is sent
		client.TrackCLIInstallEnd(false) // Will be updated if server is installed
	}()
}

// TrackServerInstall tracks server installation events
func TrackServerInstall(installType string, tlVersion string) (string, time.Time) {
	client := GetClient()
	installID := time.Now().Format("20060102150405") + "-" + client.cliID[:8]
	
	client.TrackServerInstallStart(installType, tlVersion)
	
	// Update CLI install end to indicate server was installed
	client.TrackCLIInstallEnd(true)
	
	return installID, time.Now()
}

// TrackServerInstallEnd tracks server installation completion
func TrackServerInstallEnd(installID string, startTime time.Time, wasSuccessful bool, downloadSizeGB float64) {
	client := GetClient()
	
	// Get system information
	sysInfo := GetSystemInfo()
	
	// Calculate time to install
	timeToInstall := time.Since(startTime)
	
	// Prepare machine setup info
	machineSetup := map[string]interface{}{
		"ram_gb":              sysInfo.RAM,
		"cpu_count":           sysInfo.CPUCount,
		"cpu_model":           sysInfo.CPUModel,
		"available_storage_gb": sysInfo.AvailableStorage,
		"os":                  sysInfo.OS,
		"arch":                sysInfo.Arch,
	}
	
	// Extract GPU info
	gpuCount := len(sysInfo.GPUs)
	gpuTypes := make([]string, 0, gpuCount)
	for _, gpu := range sysInfo.GPUs {
		gpuTypes = append(gpuTypes, gpu.Name)
	}
	
	// Extract mount info
	mountCount := len(sysInfo.Mounts)
	
	client.TrackServerInstallEnd(
		installID,
		timeToInstall,
		sysInfo.AvailableStorage,
		mountCount,
		gpuCount,
		gpuTypes,
		downloadSizeGB,
		wasSuccessful,
		machineSetup,
	)
}
