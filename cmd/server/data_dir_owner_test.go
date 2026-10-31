//go:build !windows

package server

import (
	"os"
	"path/filepath"
	"syscall"
	"testing"
)

func stubChown(t *testing.T) *[]string {
	t.Helper()
	var calls []string
	old := chownAsRoot
	chownAsRoot = func(path string, uid int) error {
		calls = append(calls, path)
		return nil
	}
	t.Cleanup(func() { chownAsRoot = old })
	return &calls
}

func TestRootOwnedDataDirIsHandedToTheUser(t *testing.T) {
	calls := stubChown(t)
	// "/" stands in for a data dir the installer created with sudo; the stub keeps it untouched
	if err := ownRootOwnedDir("/", 1000); err != nil {
		t.Fatal(err)
	}
	if len(*calls) != 1 || (*calls)[0] != "/" {
		t.Fatalf("want one chown of /, got %v", *calls)
	}
}

func TestDataDirOwnerIsLeftAloneWhenNotRootOwned(t *testing.T) {
	if os.Geteuid() == 0 {
		t.Skip("a temp dir created as root is root-owned")
	}
	calls := stubChown(t)
	dir := t.TempDir()
	link := filepath.Join(t.TempDir(), "link")
	if err := os.Symlink("/", link); err != nil {
		t.Fatal(err)
	}
	for _, path := range []string{dir, link} {
		if err := ownRootOwnedDir(path, os.Getuid()+1); err != nil {
			t.Fatal(err)
		}
	}
	if len(*calls) != 0 {
		t.Fatalf("a dir another user owns, or a symlink, must keep its owner; chowned %v", *calls)
	}
}

func TestDataDirStaysRootsWhenRootRanTheInstall(t *testing.T) {
	calls := stubChown(t)
	if err := ownRootOwnedDir("/", 0); err != nil {
		t.Fatal(err)
	}
	if len(*calls) != 0 {
		t.Fatalf("nothing to hand over to root, chowned %v", *calls)
	}
}

func TestMissingDataDirIsAnError(t *testing.T) {
	stubChown(t)
	if err := ownRootOwnedDir(filepath.Join(t.TempDir(), "missing"), 1000); err == nil {
		t.Fatal("want an error for a missing data dir")
	}
}

func TestRootChownKeepsGroupAndSetgid(t *testing.T) {
	if os.Geteuid() != 0 {
		t.Skip("needs root to create a root-owned dir")
	}
	dir := t.TempDir()
	if err := os.Chown(dir, 0, 4242); err != nil {
		t.Fatal(err)
	}
	if err := os.Chmod(dir, 0o775|os.ModeSetgid); err != nil {
		t.Fatal(err)
	}
	if err := ownRootOwnedDir(dir, 1000); err != nil {
		t.Fatal(err)
	}
	info, err := os.Stat(dir)
	if err != nil {
		t.Fatal(err)
	}
	st := info.Sys().(*syscall.Stat_t)
	if st.Uid != 1000 || st.Gid != 4242 || info.Mode()&os.ModeSetgid == 0 || info.Mode().Perm() != 0o775 {
		t.Fatalf("want 1000:4242 2775, got %d:%d %v", st.Uid, st.Gid, info.Mode())
	}
}
