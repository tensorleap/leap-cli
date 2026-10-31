//go:build !windows

package server

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"syscall"

	"github.com/tensorleap/helm-charts/pkg/local"
	"github.com/tensorleap/leap-cli/pkg/log"
)

// ensureDataDirOwner hands the default data dir to the user who ran leap. The installer creates it with sudo as
// root:tensorleap 2775, and a shell started before this install created the tensorleap group never carries it, so
// that user (and their AI assistant) could not write there until they logged in again. The group and the setgid bit
// stay, so other members keep their access. Linux only: elsewhere the data dir is world-writable.
func ensureDataDirOwner() {
	if runtime.GOOS != "linux" {
		return
	}
	dir := local.GetServerDataDir()
	if filepath.Clean(dir) != local.DEFAULT_DATA_DIR {
		return
	}
	if err := ownRootOwnedDir(dir, invokingUID()); err != nil {
		log.VerboseLogger.Warnf("Could not make %s owned by uid %d: %v", dir, invokingUID(), err)
	}
}

// chownAsRoot changes path's owner to uid, keeping its group; tests replace it
var chownAsRoot = func(path string, uid int) error {
	if os.Geteuid() == 0 {
		return os.Lchown(path, uid, -1)
	}
	// the install already asked for sudo, so don't prompt again
	return exec.Command("sudo", "-n", "chown", "-h", strconv.Itoa(uid), path).Run()
}

// ownRootOwnedDir gives a root-owned directory to uid. A directory another user already owns, or a symlink, is left as is.
func ownRootOwnedDir(path string, uid int) error {
	if uid == 0 {
		return nil
	}
	info, err := os.Lstat(path)
	if err != nil {
		return err
	}
	st, ok := info.Sys().(*syscall.Stat_t)
	if !info.IsDir() || !ok || st.Uid != 0 {
		return nil
	}
	return chownAsRoot(path, uid)
}
