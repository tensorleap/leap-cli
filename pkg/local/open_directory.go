package local

import (
	"fmt"
	"os/exec"
	"runtime"
)


func OpenDirectoryDefault(path string) error {
	switch runtime.GOOS {
	case "darwin":
		return exec.Command("open", path).Start()
	case "linux":
		return exec.Command("xdg-open", path).Start()
	case "windows":
		return exec.Command("explorer", path).Start()
	default:
		return fmt.Errorf("unsupported platform")
	}
}
