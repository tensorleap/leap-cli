package k3d

import (
	"context"
	"fmt"
	"log"
	"strconv"

	cliutil "github.com/k3d-io/k3d/v5/cmd/util"
	"github.com/k3d-io/k3d/v5/pkg/client"
	"github.com/k3d-io/k3d/v5/pkg/runtimes"
	k3d "github.com/k3d-io/k3d/v5/pkg/types"
)

const REGISTRY_NAME = "k3d-tensorleap-registry"

func CreateRegistry(ctx context.Context, port uint, volumes []string) error {
	if existingRegistry, _ := client.RegistryGet(ctx, runtimes.SelectedRuntime, REGISTRY_NAME); existingRegistry != nil {
		log.Println("Found existing registry!")
		return nil
	}

	reg := createRegistryConfig(port, volumes)
	_, err := client.RegistryRun(ctx, runtimes.SelectedRuntime, reg)
	return err
}

func createRegistryConfig(port uint, volumes []string) *k3d.Registry {
	exposePort, err := cliutil.ParsePortExposureSpec(strconv.FormatUint(uint64(port), 10), k3d.DefaultRegistryPort)
	if err != nil {
		log.Fatalln(err)
	}

	reg := &k3d.Registry{
		Host:         REGISTRY_NAME,
		Image:        fmt.Sprintf("%s:%s", k3d.DefaultRegistryImageRepo, k3d.DefaultRegistryImageTag),
		ExposureOpts: *exposePort,
		Network:      k3d.DefaultRuntimeNetwork,
		Volumes:      volumes,
	}

	return reg
}
