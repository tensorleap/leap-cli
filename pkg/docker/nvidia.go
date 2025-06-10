package docker

import (
	"context"

	"github.com/tensorleap/leap-cli/pkg/log"
)

// HasNvidiaRuntime checks if the docker daemon provides an 'nvidia' runtime.
func HasNvidiaRuntime() (bool, error) {
	c, err := NewClient()
	if err != nil {
		return false, err
	}
	info, err := c.Info(context.Background())
	if err != nil {
		return false, err
	}
	_, ok := info.Runtimes["nvidia"]
	return ok, nil
}

// CheckNvidiaRuntime logs whether the nvidia runtime exists.
func CheckNvidiaRuntime() {
	log.Info("Checking docker-nvidia2 driver...")
	ok, err := HasNvidiaRuntime()
	if err != nil {
		log.Warnf("Failed to check for docker-nvidia2 driver: %v", err)
		return
	}
	if !ok {
		log.Info("Missing docker-nvidia2 driver. Install nvidia-docker2 to enable GPU support.")
		return
	}
	log.Info("docker-nvidia2 driver found")
}
