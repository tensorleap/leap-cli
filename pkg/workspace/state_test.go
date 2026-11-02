package workspace

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func writeFile(t *testing.T, path, content string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}

func TestLastPushedModelIsRemembered(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, filepath.Join(dir, "model", "resnet_ls.h5"), "m")

	if err := SaveLastPushedModelPath(dir, filepath.Join(dir, "model", "resnet_ls.h5")); err != nil {
		t.Fatal(err)
	}
	if got := LastPushedModelPath(dir); got != "model/resnet_ls.h5" {
		t.Fatalf("LastPushedModelPath = %q, want model/resnet_ls.h5", got)
	}
	ignore, err := os.ReadFile(filepath.Join(dir, ".tensorleap", ".gitignore"))
	if err != nil || string(ignore) != "*\n" {
		t.Fatalf(".tensorleap/.gitignore = %q, %v; want \"*\"", ignore, err)
	}
}

func TestRememberedModelThatNoLongerExistsIsIgnored(t *testing.T) {
	dir := t.TempDir()
	model := filepath.Join(dir, "model.h5")
	writeFile(t, model, "m")
	if err := SaveLastPushedModelPath(dir, model); err != nil {
		t.Fatal(err)
	}
	if err := os.Remove(model); err != nil {
		t.Fatal(err)
	}
	if got := LastPushedModelPath(dir); got != "" {
		t.Fatalf("LastPushedModelPath = %q for a deleted file, want \"\"", got)
	}
}

func TestNothingRememberedWithoutAStateFile(t *testing.T) {
	if got := LastPushedModelPath(t.TempDir()); got != "" {
		t.Fatalf("LastPushedModelPath = %q, want \"\"", got)
	}
}

func TestCorruptStateFileIsIgnoredAndReplaced(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, filepath.Join(dir, ".tensorleap", "state.json"), "{not json")
	writeFile(t, filepath.Join(dir, "m.onnx"), "m")
	if got := LastPushedModelPath(dir); got != "" {
		t.Fatalf("LastPushedModelPath = %q, want \"\"", got)
	}
	if err := SaveLastPushedModelPath(dir, filepath.Join(dir, "m.onnx")); err != nil {
		t.Fatal(err)
	}
	if got := LastPushedModelPath(dir); got != "m.onnx" {
		t.Fatalf("LastPushedModelPath = %q, want m.onnx", got)
	}
}

func TestSavingKeepsOtherStateAndAnExistingGitignore(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, filepath.Join(dir, ".tensorleap", "state.json"), `{"other": 1}`)
	writeFile(t, filepath.Join(dir, ".tensorleap", ".gitignore"), "custom\n")
	writeFile(t, filepath.Join(dir, "m.h5"), "m")

	if err := SaveLastPushedModelPath(dir, filepath.Join(dir, "m.h5")); err != nil {
		t.Fatal(err)
	}
	var state map[string]any
	data, _ := os.ReadFile(filepath.Join(dir, ".tensorleap", "state.json"))
	if err := json.Unmarshal(data, &state); err != nil {
		t.Fatal(err)
	}
	if state["other"] != float64(1) || state["lastModelPath"] != "m.h5" {
		t.Fatalf("state.json = %s", data)
	}
	if ignore, _ := os.ReadFile(filepath.Join(dir, ".tensorleap", ".gitignore")); string(ignore) != "custom\n" {
		t.Fatalf("an existing .gitignore was rewritten: %q", ignore)
	}
}

func TestModelOutsideTheProjectFolderIsNotRemembered(t *testing.T) {
	dir := t.TempDir()
	outside := filepath.Join(t.TempDir(), "m.h5")
	writeFile(t, outside, "m")
	if err := SaveLastPushedModelPath(dir, outside); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(dir, ".tensorleap")); !os.IsNotExist(err) {
		t.Fatalf("a model outside the folder created .tensorleap/: %v", err)
	}
}
