package workspace

import (
	"os"
	"strings"
	"testing"
)

func TestSetProjectIdIfMissing(t *testing.T) {
	t.Run("no leap.yaml", func(t *testing.T) {
		t.Chdir(t.TempDir())
		written, existing, err := SetProjectIdIfMissing("p-new")
		if err != nil || written || existing != "" {
			t.Fatalf("got written=%v existing=%q err=%v", written, existing, err)
		}
		if _, err := os.Stat(CONFIG_FILE_NAME); !os.IsNotExist(err) {
			t.Fatalf("leap.yaml was created")
		}
	})

	t.Run("leap.yaml without a project", func(t *testing.T) {
		t.Chdir(t.TempDir())
		writeConfig(t, "projectId: \nsecretId: \nentryFile: leap_integration.py\ninclude:\n  - leap_integration.py\n")
		written, existing, err := SetProjectIdIfMissing("p-new")
		if err != nil || !written || existing != "" {
			t.Fatalf("got written=%v existing=%q err=%v", written, existing, err)
		}
		content := readConfig(t)
		if !strings.Contains(content, "projectId: p-new") || !strings.Contains(content, "- leap_integration.py") {
			t.Fatalf("unexpected leap.yaml:\n%s", content)
		}
	})

	t.Run("leap.yaml with another project", func(t *testing.T) {
		t.Chdir(t.TempDir())
		original := "projectId: p-old\nsecretId: \nentryFile: leap_integration.py\ninclude:\n  - leap_integration.py\n"
		writeConfig(t, original)
		written, existing, err := SetProjectIdIfMissing("p-new")
		if err != nil || written || existing != "p-old" {
			t.Fatalf("got written=%v existing=%q err=%v", written, existing, err)
		}
		if content := readConfig(t); content != original {
			t.Fatalf("leap.yaml changed:\n%s", content)
		}
	})
}

func writeConfig(t *testing.T, content string) {
	t.Helper()
	if err := os.WriteFile(CONFIG_FILE_NAME, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
}

func readConfig(t *testing.T) string {
	t.Helper()
	content, err := os.ReadFile(CONFIG_FILE_NAME)
	if err != nil {
		t.Fatal(err)
	}
	return string(content)
}
