package secret

import (
	"context"
	"fmt"

	"github.com/tensorleap/leap-cli/pkg/api"
	"github.com/tensorleap/leap-cli/pkg/entity"
	"github.com/tensorleap/leap-cli/pkg/log"
	"github.com/tensorleap/leap-cli/pkg/tensorleapapi"
)

type SecretEntity = tensorleapapi.SecretManager

var SecretEntityDesc = entity.NewEntityDescriptor[SecretEntity](
	"Secret",
	"Secrets",
	func(s *SecretEntity) string { return s.Name },
	func(s *SecretEntity) string { return s.GetCid() },
)

func GetSecretList(ctx context.Context) ([]SecretEntity, error) {
	list, _, err := api.ApiClient.GetSecretManagerList(ctx).Execute()
	return list.Results, err
}

func CreateSecretManager(ctx context.Context, name, key string) (*SecretEntity, error) {
	log.Infof("Creating secret manager: %s", name)
	params := tensorleapapi.NewAddSecretManagerParams(name, key)
	_, _, err := api.ApiClient.AddSecretManager(ctx).AddSecretManagerParams(*params).Execute()
	if err != nil {
		return nil, err
	}
	list, err := GetSecretList(ctx)
	if err != nil {
		return nil, err
	}
	newSecret, err := entity.GetEntityByDisplayName(name, list, SecretEntityDesc)
	if err != nil {
		return nil, fmt.Errorf("failed to add secret manager: %w", err)
	}
	entity.InfoCreation(newSecret, SecretEntityDesc)
	return newSecret, nil

}

func DeleteSecretManager(ctx context.Context, secret *SecretEntity) error {
	log.Infof("Deleting secret manager: %s, id: %s", secret.Name, secret.GetCid())
	params := tensorleapapi.NewTrashSecretManagerParams(secret.GetCid())
	_, _, err := api.ApiClient.TrashSecretManager(ctx).TrashSecretManagerParams(*params).Execute()
	if err != nil {
		return err
	}
	entity.InfoDeletion(secret, SecretEntityDesc)
	return nil
}
