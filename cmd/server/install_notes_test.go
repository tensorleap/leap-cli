//go:build !windows

package server

import (
	"os"
	"path/filepath"
	"syscall"
	"testing"
)

func TestInstallNotesAreCreatedOwnedByTheUserAndNeverOverwritten(t *testing.T) {
	path := filepath.Join(t.TempDir(), installNotesFile)
	if err := createOwnedFile(path, os.Getuid()); err != nil {
		t.Fatal(err)
	}
	info, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	if info.Mode().Perm()&0o600 != 0o600 {
		t.Fatalf("the owner must read and write it, mode %v", info.Mode())
	}
	if st, ok := info.Sys().(*syscall.Stat_t); ok && int(st.Uid) != os.Getuid() {
		t.Fatalf("owned by uid %d, want %d", st.Uid, os.Getuid())
	}

	if err := os.WriteFile(path, []byte("# notes\n"), 0o664); err != nil {
		t.Fatal(err)
	}
	if err := createOwnedFile(path, os.Getuid()); err != nil {
		t.Fatal(err)
	}
	if b, _ := os.ReadFile(path); string(b) != "# notes\n" {
		t.Fatalf("a later install must keep the notes, got %q", b)
	}
}

func TestInstallNotesNeedTheDataDir(t *testing.T) {
	if err := createOwnedFile(filepath.Join(t.TempDir(), "missing", installNotesFile), os.Getuid()); err == nil {
		t.Fatal("no data dir, no notes; and no sudo attempt for it")
	}
}
