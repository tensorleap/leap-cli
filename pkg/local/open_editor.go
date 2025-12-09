package local

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/tensorleap/leap-cli/pkg/log"
)

// OpenDirectorySmart tries to open a directory in this order:
// 1. PyCharm (or charm)
// 2. VS Code (code)
// 3. Cursor (cursor)
// 4. vim
// 5. vi
// 6. OS file manager (Finder / Explorer / xdg-open)
func OpenDirectorySmart(dir string) error {
	// Normalize path (optional, but nice to have)
	absDir, err := filepath.Abs(dir)
	if err != nil {
		absDir = dir
	}

	// 1) PyCharm
	if openWithPyCharm(absDir) {
		return nil
	}

	// 2) VS Code
	if openWithVSCode(absDir) {
		return nil
	}

	// 3) Cursor
	if openWithCursor(absDir) {
		return nil
	}

	// 4) vim
	if tryEditorByName("vim", absDir) {
		return nil
	}

	// 5) vi
	if tryEditorByName("vi", absDir) {
		return nil
	}

	// 6) Fallback: OS file manager
	return openWithFileManager(absDir)
}

// ---------- Generic helpers ----------

// tryEditorByName tries to run an editor by command name if it's in PATH.
// Returns true if it started successfully.
func tryEditorByName(cmdName string, args ...string) bool {
	_, err := exec.LookPath(cmdName)
	if err != nil {
		return false
	}
	cmd := exec.Command(cmdName, args...)
	_ = cmd.Start() // we don't wait; fire-and-forget
	return true
}

// tryEditorByPath tries to run an editor at a specific full path if it exists.
// Returns true if it started successfully.
func tryEditorByPath(fullPath string, args ...string) bool {
	if fullPath == "" {
		return false
	}
	if st, err := os.Stat(fullPath); err != nil || st.IsDir() {
		return false
	}
	cmd := exec.Command(fullPath, args...)
	_ = cmd.Start()
	return true
}

// tryPathCandidates tries each candidate path until one succeeds.
// Returns true if any candidate started successfully.
func tryPathCandidates(candidates []string, args ...string) bool {
	for _, c := range candidates {
		if tryEditorByPath(c, args...) {
			return true
		}
	}
	return false
}

// openWithFileManager opens a directory in the OS file manager.
func openWithFileManager(dir string) error {
	switch runtime.GOOS {
	case "darwin":
		return exec.Command("open", dir).Start()
	case "linux":
		return exec.Command("xdg-open", dir).Start()
	case "windows":
		return exec.Command("explorer", dir).Start()
	default:
		return fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}
}

// homeDir returns the current user's home directory (best-effort).
func homeDir() string {
	if h, err := os.UserHomeDir(); err == nil {
		return h
	}
	// fallback envs
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	if h := os.Getenv("USERPROFILE"); h != "" {
		return h
	}
	return ""
}

// ---------- PyCharm ----------

func openWithPyCharm(dir string) bool {
	// 1) CLI in PATH: pycharm / charm
	if tryEditorByName("pycharm", dir) {
		return true
	}
	if tryEditorByName("charm", dir) {
		return true
	}

	var candidates []string

	switch runtime.GOOS {
	case "darwin":
		candidates = []string{
			"/Applications/PyCharm.app/Contents/MacOS/pycharm",
		}

	case "linux":
		home := homeDir()
		if home != "" {
			pattern := filepath.Join(home, ".local", "share", "JetBrains", "Toolbox", "apps", "PyCharm", "*", "bin", "pycharm.sh")
			candidates, _ = filepath.Glob(pattern)
		}

	case "windows":
		progFiles := os.Getenv("ProgramFiles")
		if progFiles != "" {
			pattern := filepath.Join(progFiles, "JetBrains", "PyCharm*", "bin", "pycharm64.exe")
			candidates, _ = filepath.Glob(pattern)
		}
		localAppData := os.Getenv("LOCALAPPDATA")
		if localAppData != "" {
			pattern := filepath.Join(localAppData, "JetBrains", "Toolbox", "apps", "PyCharm", "*", "bin", "pycharm64.exe")
			toolboxCandidates, err := filepath.Glob(pattern)
			if err != nil {
				log.Warnf("failed to get toolbox candidates: %v", err)
			} else {
				candidates = append(candidates, toolboxCandidates...)
			}
		}
	}

	return tryPathCandidates(candidates, dir)
}

// ---------- VS Code ----------

func openWithVSCode(dir string) bool {
	// 1) CLI in PATH
	if tryEditorByName("code", dir) {
		return true
	}

	var candidates []string

	switch runtime.GOOS {
	case "darwin":
		candidates = []string{
			"/Applications/Visual Studio Code.app/Contents/Resources/app/bin/code",
		}
	case "linux":
		candidates = []string{
			"/usr/bin/code",
		}
	case "windows":
		localAppData := os.Getenv("LOCALAPPDATA")
		if localAppData != "" {
			candidates = []string{
				filepath.Join(localAppData, "Programs", "Microsoft VS Code", "Code.exe"),
			}
		}
	}

	return tryPathCandidates(candidates, dir)
}

// ---------- Cursor ----------

func openWithCursor(dir string) bool {
	// 1) CLI in PATH
	if tryEditorByName("cursor", dir) {
		return true
	}

	var candidates []string

	switch runtime.GOOS {
	case "darwin":
		candidates = []string{
			"/Applications/Cursor.app/Contents/MacOS/Cursor",
		}
	case "windows":
		localAppData := os.Getenv("LOCALAPPDATA")
		if localAppData != "" {
			candidates = []string{
				filepath.Join(localAppData, "Programs", "Cursor", "Cursor.exe"),
			}
		}
		// Linux: usually "cursor" would be in PATH if installed
	}

	return tryPathCandidates(candidates, dir)
}
