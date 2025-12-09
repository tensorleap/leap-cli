package local

import (
	"fmt"
	"os/exec"
	"runtime"
)

func OpenLink(link string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", link)
	case "linux":
		cmd = exec.Command("xdg-open", link)
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", link)
	default:
		return fmt.Errorf("unsupported platform")
	}

	return cmd.Start()
}

