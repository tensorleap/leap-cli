package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func CreateDockerCli() (*client.Client, error) {
	return client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
}

func FindContainerIDByName(ctx context.Context, cli *client.Client, containerName string) (string, error) {
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{All: true})
	if err != nil {
		return "", err
	}

	for _, container := range containers {
		for _, name := range container.Names {
			if name == "/"+containerName {
				return container.ID, nil
			}
		}
	}

	return "", fmt.Errorf("container '%s' not found", containerName)
}

// RemoveContainer removes the specified container by its ID.
func RemoveContainer(ctx context.Context, cli *client.Client, containerID string) error {
	err := cli.ContainerRemove(ctx, containerID, types.ContainerRemoveOptions{Force: true})
	if err != nil {
		return err
	}
	return nil
}
