package secret

import (
	"context"
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/tensorleap/leap-cli/pkg/entity"
	"github.com/tensorleap/leap-cli/pkg/log"
	"github.com/tensorleap/leap-cli/pkg/workspace"
)

func SelectOrCreateSecret(ctx context.Context, secrets []SecretEntity, askForNewFirst bool) (*SecretEntity, bool, error) {
	create := func() (*SecretEntity, error) {
		var name string
		var key string
		var keyPath string

		err := AskForNewSecret(ctx, secrets, &name, &key, keyPath)
		if err != nil {
			return nil, err

		}

		secret, err := CreateSecretManager(ctx, name, key)
		if err != nil {
			return nil, err
		}
		return secret, nil
	}
	return entity.SelectEntityOrCreateOne(secrets, create, askForNewFirst, SecretEntityDesc)
}

func AskForNewSecret(ctx context.Context, secrets []SecretEntity, name, key *string, keyPath string) error {

	existingNames := entity.GetNames(secrets, SecretEntityDesc)
	validator := entity.CreateUniqueNameValidator(existingNames)

	err := validator(*name)
	isNameNotValid := err != nil
	if isNameNotValid {
		if len(*name) > 0 {
			log.Warn(err.Error())
		}

		*name, err = entity.AskForName(existingNames, *name, SecretEntityDesc)
		if err != nil {
			return err
		}
	}

	if len(*key) == 0 && len(keyPath) == 0 {
		err = survey.AskOne(&survey.Input{
			Message: "Enter secret key path:",
		}, &keyPath)

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

func AskIfUseSecret() (bool, error) {
	var useSecret bool
	err := survey.AskOne(&survey.Confirm{
		Message: "Do you want to use a secret?",
	}, &useSecret)
	if err != nil {
		return false, err
	}
	return useSecret, nil
}

func CreateOrSelectSecretIfInUse(ctx context.Context) (selectedSecret *SecretEntity, wasCreateNew bool, err error) {
	useSecret, err := AskIfUseSecret()

	if err != nil {
		return nil, false, err
	}

	if !useSecret {
		return nil, false, nil
	}

	secrets, err := GetSecretList(ctx)
	if err != nil {
		return nil, false, err
	}

	selectedSecret, wasCreateNew, err = SelectOrCreateSecret(ctx, secrets, true)
	if err != nil {
		return nil, false, err
	}
	return selectedSecret, wasCreateNew, nil
}

func CreateOrSelectIfSecretNotFound(ctx context.Context, secretId string) (selectedSecret *SecretEntity, wasValid, wasCreateNew bool, err error) {
	secrets, err := GetSecretList(ctx)
	if err != nil {
		return nil, false, false, err
	}
	selectedSecret, _ = entity.GetEntityById(secretId, secrets, SecretEntityDesc)
	if selectedSecret != nil {
		return selectedSecret, true, false, nil
	}
	log.Warnf("Secret with id %s not found", secretId)

	isUsingSecret, err := AskIfUseSecret()
	if err != nil {
		return nil, false, false, err
	}
	if !isUsingSecret {
		return nil, false, false, nil
	}

	selectedSecret, wasCreateNew, err = SelectOrCreateSecret(ctx, secrets, true)
	if err != nil {
		return nil, false, false, err
	}
	return selectedSecret, false, wasCreateNew, nil
}

func SyncSecretIdFromFlagAndConfig(ctx context.Context, secretId string, workspaceConfig *workspace.WorkspaceConfig) (string, error) {
	if len(secretId) == 0 {
		secretId = workspaceConfig.SecretId
	}

	if len(secretId) > 0 {
		selectedSecrete, _, _, err := CreateOrSelectIfSecretNotFound(ctx, secretId)
		if err != nil {
			return "", err
		}
		secretId = selectedSecrete.GetCid()
		if workspaceConfig.SecretId != secretId {
			log.Infof("Updating leap.yaml secret id to %s", secretId)
			workspaceConfig.SecretId = secretId
			err = workspace.SetWorkspaceConfig(workspaceConfig, ".")
			if err != nil {
				return "", err
			}
		}
	}
	return secretId, nil
}

func SetSecret(ctx context.Context, flagSecretId string) error {
	workspaceConfig, err := workspace.GetWorkspaceConfig()
	if err != nil {
		return err
	}
	secrets, err := GetSecretList(ctx)
	if err != nil {
		return err
	}
	var selectedSecret *SecretEntity
	if len(flagSecretId) == 0 {
		if workspaceConfig.SecretId != "" {
			currentSecret, _ := entity.GetEntityById(workspaceConfig.SecretId, secrets, SecretEntityDesc)
			if currentSecret != nil {
				var confirmed bool
				err = survey.AskOne(&survey.Confirm{
					Message: fmt.Sprintf("You are already using secret: %s, Do you want to change secret", currentSecret.GetName()),
					Default: true,
				}, &confirmed)

				if err != nil {
					return err
				}
				if !confirmed {
					return nil
				}
			}
		}

		selectedSecret, _, err = SelectOrCreateSecret(ctx, secrets, true)
		if err != nil {
			return err
		}
	} else {
		selectedSecret, _, _, err = CreateOrSelectIfSecretNotFound(ctx, flagSecretId)
		if err != nil {
			return err
		}
	}

	workspaceConfig.SecretId = selectedSecret.GetCid()
	log.Infof("Set leap.yaml with secret ID: %s", workspaceConfig.SecretId)
	err = workspace.SetWorkspaceConfig(workspaceConfig, ".")

	return err
}
