package workspace

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// The project folder's local state lives in .tensorleap/, next to leap.yaml. It is per-machine memory
// (never pushed: the code bundle always skips .tensorleap/, and the folder ignores itself in git).
const (
	LocalStateDir  = ".tensorleap"
	localStateFile = "state.json"
	lastModelKey   = "lastModelPath"
)

func localStatePath(projectDir string) string {
	return filepath.Join(projectDir, LocalStateDir, localStateFile)
}

// readLocalState returns every key of state.json, so a write keeps the keys it doesn't know about.
// A missing or unreadable file is an empty state: it only holds conveniences.
func readLocalState(projectDir string) map[string]json.RawMessage {
	state := map[string]json.RawMessage{}
	data, err := os.ReadFile(localStatePath(projectDir))
	if err != nil {
		return state
	}
	if err := json.Unmarshal(data, &state); err != nil || state == nil {
		return map[string]json.RawMessage{}
	}
	return state
}

// LastPushedModelPath is the model file of the last successful push from this folder, relative to it
// (e.g. "model/resnet_ls.h5"), or "" when none is remembered or the file no longer exists.
func LastPushedModelPath(projectDir string) string {
	raw, ok := readLocalState(projectDir)[lastModelKey]
	if !ok {
		return ""
	}
	var rel string
	if err := json.Unmarshal(raw, &rel); err != nil || rel == "" {
		return ""
	}
	if info, err := os.Stat(filepath.Join(projectDir, rel)); err != nil || info.IsDir() {
		return ""
	}
	return rel
}

// SaveLastPushedModelPath remembers modelPath (absolute, or relative to the working directory) as the
// model of the last successful push from projectDir. A model outside projectDir isn't remembered: the
// picker only lists files inside it.
func SaveLastPushedModelPath(projectDir, modelPath string) error {
	rel, ok := relativeInside(projectDir, modelPath)
	if !ok {
		return nil
	}
	if err := ensureLocalStateDir(projectDir); err != nil {
		return err
	}
	state := readLocalState(projectDir)
	encoded, err := json.Marshal(filepath.ToSlash(rel))
	if err != nil {
		return err
	}
	state[lastModelKey] = encoded
	data, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return err
	}
	return writeFileAtomic(localStatePath(projectDir), append(data, '\n'))
}

func relativeInside(projectDir, path string) (string, bool) {
	absDir, err := filepath.Abs(projectDir)
	if err != nil {
		return "", false
	}
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", false
	}
	rel, err := filepath.Rel(absDir, absPath)
	if err != nil || rel == "." || rel == ".." || strings.HasPrefix(rel, ".."+string(filepath.Separator)) {
		return "", false
	}
	return rel, true
}

// ensureLocalStateDir creates .tensorleap/ with a .gitignore of "*", so nothing in it is committed.
// An existing .gitignore is left as it is.
func ensureLocalStateDir(projectDir string) error {
	dir := filepath.Join(projectDir, LocalStateDir)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("create %s: %w", dir, err)
	}
	gitignore := filepath.Join(dir, ".gitignore")
	f, err := os.OpenFile(gitignore, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o644)
	if errors.Is(err, fs.ErrExist) {
		return nil
	}
	if err != nil {
		return fmt.Errorf("create %s: %w", gitignore, err)
	}
	if _, err := f.WriteString("*\n"); err != nil {
		_ = f.Close()
		return fmt.Errorf("write %s: %w", gitignore, err)
	}
	return f.Close()
}

func writeFileAtomic(path string, data []byte) error {
	tmp, err := os.CreateTemp(filepath.Dir(path), filepath.Base(path)+".*.tmp")
	if err != nil {
		return err
	}
	if _, err := tmp.Write(data); err != nil {
		_ = tmp.Close()
		_ = os.Remove(tmp.Name())
		return err
	}
	if err := tmp.Close(); err != nil {
		_ = os.Remove(tmp.Name())
		return err
	}
	if err := os.Rename(tmp.Name(), path); err != nil {
		_ = os.Remove(tmp.Name())
		return err
	}
	return nil
}
