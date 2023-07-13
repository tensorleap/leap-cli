package local

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/tensorleap/cli-go/pkg/k3d"
)

func ValidateStandaloneDir() error {
	_, err := os.Stat(STANDALONE_DIR)
	if os.IsNotExist(err) {
		return errors.New(
			fmt.Sprintf("Not found data directory(%s) on this machine, Please make sure to install before upgrade", STANDALONE_DIR),
		)
	}
	return err
}

func GetTensorleapCluster(ctx context.Context) (*k3d.Cluster, error) {
	cluster, err := k3d.GetCluster(ctx)
	if err != nil {
		return nil, err
	}
	if cluster == nil {
		return nil, errors.New(
			fmt.Sprintf("Not found local Cluster(%s) on this machine, Please make sure to install before upgrade", STANDALONE_DIR),
		)
	}
	return cluster, nil
}
