package secret

import (
	"context"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/tensorleap/leap-cli/pkg/entity"
	"github.com/tensorleap/leap-cli/pkg/log"
)

func SelectOrCreateSecret(ctx context.Context, secrets []SecretEntity, askForNewProjectFirst bool) (*SecretEntity, bool, error) {
	create := func() (*SecretEntity, error) {
		var name string
		var key string
		var keyPath string

		err := AskForSecret(ctx, secrets, &name, &key, keyPath)
		if err != nil {
			return nil, err

		}

		secret, err := CreateSecretManager(ctx, name, key)
		if err != nil {
			return nil, err
		}
		return secret, nil
	}
	return entity.SelectEntityOrCreateOne(secrets, create, askForNewProjectFirst, SecretEntityDesc)
}

func AskForSecret(ctx context.Context, secrets []SecretEntity, name, key *string, keyPath string) error {

	existingNames := entity.GetNames(secrets, SecretEntityDesc)
	validator := entity.CreateUniqueNameValidator(existingNames)

	err := validator(*name)
	isNameNotValid := err != nil
	if isNameNotValid {
		if len(*name) > 0 {
			log.Warnf(err.Error())
		}

		*name, err = entity.AskForName(existingNames, SecretEntityDesc)
		if err != nil {
			return err
		}
	}

	if len(*key) == 0 && len(keyPath) == 0 {
		err = survey.AskOne(&survey.Input{
			Message: "Enter secret key path:",
		}, keyPath)

		if err != nil {
			return err
		}
	}

	if len(keyPath) > 0 {
		keyBytes, err := os.ReadFile(keyPath)
		if err != nil {
			return err
		}
		*key = string(keyBytes)
	}

	return nil
}
