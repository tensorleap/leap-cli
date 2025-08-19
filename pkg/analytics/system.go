package analytics

import (
	"bufio"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"syscall"
)

// SystemInfo contains information about the machine
type SystemInfo struct {
	RAM            int64   `json:"ram_gb"`
	CPUCount       int     `json:"cpu_count"`
	CPUModel       string  `json:"cpu_model"`
	AvailableStorage int64 `json:"available_storage_gb"`
	OS             string  `json:"os"`
	Arch           string  `json:"arch"`
	GPUs           []GPU   `json:"gpus"`
	Mounts         []Mount `json:"mounts"`
}

// GPU contains information about a GPU
type GPU struct {
	Name   string `json:"name"`
	Memory string `json:"memory"`
	Type   string `json:"type"`
}

// Mount contains information about a mount point
type Mount struct {
	Device     string `json:"device"`
	MountPoint string `json:"mount_point"`
	FileSystem string `json:"file_system"`
}

// GetSystemInfo gathers system information for analytics
func GetSystemInfo() *SystemInfo {
	info := &SystemInfo{
		OS:   runtime.GOOS,
		Arch: runtime.GOARCH,
	}

	// Get RAM info
	info.RAM = getRAMInfo()
	
	// Get CPU info
	info.CPUCount = runtime.NumCPU()
	info.CPUModel = getCPUModel()
	
	// Get storage info
	info.AvailableStorage = getAvailableStorage()
	
	// Get GPU info
	info.GPUs = getGPUInfo()
	
	// Get mount info
	info.Mounts = getMountInfo()

	return info
}

// getRAMInfo gets RAM information in GB
func getRAMInfo() int64 {
	if runtime.GOOS == "linux" {
		// Read /proc/meminfo on Linux
		file, err := os.Open("/proc/meminfo")
		if err != nil {
			return 0
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			if strings.HasPrefix(line, "MemTotal:") {
				parts := strings.Fields(line)
				if len(parts) >= 2 {
					if kb, err := strconv.ParseInt(parts[1], 10, 64); err == nil {
						return kb / 1024 / 1024 // Convert KB to GB
					}
				}
			}
		}
	} else if runtime.GOOS == "darwin" {
		// Use sysctl on macOS
		cmd := exec.Command("sysctl", "-n", "hw.memsize")
		if output, err := cmd.Output(); err == nil {
			if bytes, err := strconv.ParseInt(strings.TrimSpace(string(output)), 10, 64); err == nil {
				return bytes / 1024 / 1024 / 1024 // Convert bytes to GB
			}
		}
	} else if runtime.GOOS == "windows" {
		// Use wmic on Windows
		cmd := exec.Command("wmic", "computersystem", "get", "TotalPhysicalMemory", "/value")
		if output, err := cmd.Output(); err == nil {
			lines := strings.Split(string(output), "\n")
			for _, line := range lines {
				if strings.Contains(line, "TotalPhysicalMemory=") {
					parts := strings.Split(line, "=")
					if len(parts) >= 2 {
						if bytes, err := strconv.ParseInt(strings.TrimSpace(parts[1]), 10, 64); err == nil {
							return bytes / 1024 / 1024 / 1024 // Convert bytes to GB
						}
					}
				}
			}
		}
	}
	return 0
}

// getCPUModel gets the CPU model name
func getCPUModel() string {
	if runtime.GOOS == "linux" {
		// Read /proc/cpuinfo on Linux
		file, err := os.Open("/proc/cpuinfo")
		if err != nil {
			return "unknown"
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			if strings.HasPrefix(line, "model name") {
				parts := strings.Split(line, ":")
				if len(parts) >= 2 {
					return strings.TrimSpace(parts[1])
				}
			}
		}
	} else if runtime.GOOS == "darwin" {
		// Use sysctl on macOS
		cmd := exec.Command("sysctl", "-n", "machdep.cpu.brand_string")
		if output, err := cmd.Output(); err == nil {
			return strings.TrimSpace(string(output))
		}
	} else if runtime.GOOS == "windows" {
		// Use wmic on Windows
		cmd := exec.Command("wmic", "cpu", "get", "name", "/value")
		if output, err := cmd.Output(); err == nil {
			lines := strings.Split(string(output), "\n")
			for _, line := range lines {
				if strings.Contains(line, "Name=") {
					parts := strings.Split(line, "=")
					if len(parts) >= 2 {
						return strings.TrimSpace(parts[1])
					}
				}
			}
		}
	}
	return "unknown"
}

// getAvailableStorage gets available storage in GB
func getAvailableStorage() int64 {
	var stat syscall.Statfs_t
	wd, err := os.Getwd()
	if err != nil {
		return 0
	}
	
	if err := syscall.Statfs(wd, &stat); err != nil {
		return 0
	}
	
	// Calculate available space in GB
	available := stat.Bavail * uint64(stat.Bsize)
	return int64(available / 1024 / 1024 / 1024)
}

// getGPUInfo gets GPU information
func getGPUInfo() []GPU {
	var gpus []GPU
	
	if runtime.GOOS == "linux" {
		// Try to get GPU info using lspci
		cmd := exec.Command("lspci", "-v")
		if output, err := cmd.Output(); err == nil {
			lines := strings.Split(string(output), "\n")
			var currentGPU GPU
			
			for _, line := range lines {
				if strings.Contains(line, "VGA") || strings.Contains(line, "3D") {
					if currentGPU.Name != "" {
						gpus = append(gpus, currentGPU)
					}
					currentGPU = GPU{}
					parts := strings.Fields(line)
					if len(parts) >= 4 {
						currentGPU.Name = strings.Join(parts[3:], " ")
						currentGPU.Type = "GPU"
					}
				} else if strings.Contains(line, "Memory:") && currentGPU.Name != "" {
					parts := strings.Fields(line)
					if len(parts) >= 2 {
						currentGPU.Memory = strings.TrimSpace(parts[1])
					}
				}
			}
			if currentGPU.Name != "" {
				gpus = append(gpus, currentGPU)
			}
		}
	} else if runtime.GOOS == "darwin" {
		// Use system_profiler on macOS
		cmd := exec.Command("system_profiler", "SPDisplaysDataType")
		if output, err := cmd.Output(); err == nil {
			lines := strings.Split(string(output), "\n")
			var currentGPU GPU
			
			for _, line := range lines {
				line = strings.TrimSpace(line)
				if strings.Contains(line, "Chipset Model:") {
					if currentGPU.Name != "" {
						gpus = append(gpus, currentGPU)
					}
					currentGPU = GPU{}
					parts := strings.Split(line, ":")
					if len(parts) >= 2 {
						currentGPU.Name = strings.TrimSpace(parts[1])
						currentGPU.Type = "GPU"
					}
				} else if strings.Contains(line, "VRAM:") && currentGPU.Name != "" {
					parts := strings.Split(line, ":")
					if len(parts) >= 2 {
						currentGPU.Memory = strings.TrimSpace(parts[1])
					}
				}
			}
			if currentGPU.Name != "" {
				gpus = append(gpus, currentGPU)
			}
		}
	} else if runtime.GOOS == "windows" {
		// Use wmic on Windows
		cmd := exec.Command("wmic", "path", "win32_VideoController", "get", "name,adapterram", "/value")
		if output, err := cmd.Output(); err == nil {
			lines := strings.Split(string(output), "\n")
			var currentGPU GPU
			
			for _, line := range lines {
				if strings.Contains(line, "Name=") {
					if currentGPU.Name != "" {
						gpus = append(gpus, currentGPU)
					}
					currentGPU = GPU{}
					parts := strings.Split(line, "=")
					if len(parts) >= 2 {
						currentGPU.Name = strings.TrimSpace(parts[1])
						currentGPU.Type = "GPU"
					}
				} else if strings.Contains(line, "AdapterRAM=") && currentGPU.Name != "" {
					parts := strings.Split(line, "=")
					if len(parts) >= 2 {
						if ram, err := strconv.ParseInt(strings.TrimSpace(parts[1]), 10, 64); err == nil {
							currentGPU.Memory = strconv.FormatInt(ram/1024/1024, 10) + " MB"
						}
					}
				}
			}
			if currentGPU.Name != "" {
				gpus = append(gpus, currentGPU)
			}
		}
	}
	
	return gpus
}

// getMountInfo gets mount point information
func getMountInfo() []Mount {
	var mounts []Mount
	
	if runtime.GOOS == "linux" {
		// Read /proc/mounts on Linux
		file, err := os.Open("/proc/mounts")
		if err != nil {
			return mounts
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			parts := strings.Fields(line)
			if len(parts) >= 3 {
				mount := Mount{
					Device:     parts[0],
					MountPoint: parts[1],
					FileSystem: parts[2],
				}
				mounts = append(mounts, mount)
			}
		}
	} else if runtime.GOOS == "darwin" {
		// Use mount on macOS
		cmd := exec.Command("mount")
		if output, err := cmd.Output(); err == nil {
			lines := strings.Split(string(output), "\n")
			for _, line := range lines {
				parts := strings.Fields(line)
				if len(parts) >= 3 {
					mount := Mount{
						Device:     parts[0],
						MountPoint: parts[2],
						FileSystem: parts[3],
					}
					mounts = append(mounts, mount)
				}
			}
		}
	} else if runtime.GOOS == "windows" {
		// Use wmic on Windows
		cmd := exec.Command("wmic", "logicaldisk", "get", "deviceid,size,freespace", "/value")
		if output, err := cmd.Output(); err == nil {
			lines := strings.Split(string(output), "\n")
			var currentMount Mount
			
			for _, line := range lines {
				if strings.Contains(line, "DeviceID=") {
					if currentMount.Device != "" {
						mounts = append(mounts, currentMount)
					}
					currentMount = Mount{}
					parts := strings.Split(line, "=")
					if len(parts) >= 2 {
						currentMount.Device = strings.TrimSpace(parts[1])
						currentMount.MountPoint = strings.TrimSpace(parts[1]) + "\\"
						currentMount.FileSystem = "NTFS" // Default for Windows
					}
				}
			}
			if currentMount.Device != "" {
				mounts = append(mounts, currentMount)
			}
		}
	}
	
	return mounts
}
