package server

import (
	"errors"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"

	"github.com/tensorleap/helm-charts/pkg/local"
	"github.com/tensorleap/leap-cli/pkg/log"
)

// installNotesFile is the install skill's record of this machine's setup, at the data dir root
const installNotesFile = "install-notes.md"

// ensureInstallNotes creates <data-dir>/install-notes.md owned by the user running leap. On Linux the data dir
// belongs to the tensorleap group, and a user this install just added to it has a shell without the group
// until they log in again; owning the file lets them (and their AI assistant) write it from that shell.
// The data dir itself is handed to that user first, so they can also create and replace files next to it.
func ensureInstallNotes() {
	ensureDataDirOwner()
	if err := createOwnedFile(filepath.Join(local.GetServerDataDir(), installNotesFile), invokingUID()); err != nil {
		log.VerboseLogger.Warnf("Could not create %s: %v", installNotesFile, err)
	}
}

// createOwnedFile makes path exist as a file owned by uid; an existing file is left as is
func createOwnedFile(path string, uid int) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o664)
	switch {
	case err == nil:
		_ = f.Close()
		if runtime.GOOS != "windows" && os.Geteuid() != uid {
			return os.Chown(path, uid, -1)
		}
		return nil
	case errors.Is(err, fs.ErrExist):
		return nil
	case errors.Is(err, fs.ErrPermission) && runtime.GOOS != "windows":
		// this process may not carry the group; the install already asked for sudo, so don't prompt again
		return exec.Command("sudo", "-n", "install", "-m", "0664", "-o", strconv.Itoa(uid), "/dev/null", path).Run()
	}
	return err
}

// invokingUID is the user who ran leap, also under sudo
func invokingUID() int {
	if os.Geteuid() == 0 {
		if uid, err := strconv.Atoi(os.Getenv("SUDO_UID")); err == nil {
			return uid
		}
	}
	return os.Getuid()
}
