//go:build windows
// +build windows

package server

func init() {
	isNotSupported = true
}
