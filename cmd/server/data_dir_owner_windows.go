//go:build windows

package server

// ensureDataDirOwner is a no-op on Windows, where the server commands are not supported
func ensureDataDirOwner() {}
