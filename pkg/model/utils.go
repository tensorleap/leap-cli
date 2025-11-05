package model

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/tensorleap/leap-cli/pkg/log"
	tlApi "github.com/tensorleap/leap-cli/pkg/tensorleapapi"
	"k8s.io/utils/strings/slices"
)

var MODEL_TYPES = []string{
	string(tlApi.IMPORTMODELTYPE_JSON_TF2),
	string(tlApi.IMPORTMODELTYPE_ONNX),
	string(tlApi.IMPORTMODELTYPE_PB_TF2),
	string(tlApi.IMPORTMODELTYPE_H5_TF2),
}

func GetDefaultMessageFromModelPath(filePath string) string {
	fileName := filepath.Base(filePath)
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

func InitMessage(msg *string, defaultMsg string) error {
	if len(*msg) > 0 {
		return nil
	}
	validate := func(val interface{}) error {
		str, ok := val.(string)
		if !ok {
			return fmt.Errorf("expected string, got %T", val)
		}
		str = strings.TrimSpace(str)
		if len(str) == 0 {
			return fmt.Errorf("message cannot be empty")
		}
		return nil
	}
	msgPrompt := &survey.Input{
		Message: "Enter model version message",
		Help:    "This message will be displayed in the version control page",
		Default: defaultMsg,
	}
	err := survey.AskOne(msgPrompt, msg, survey.WithValidator(validate))
	if err != nil {
		return err
	}
	return nil
}

func SelectModelType(modelType *string, modelPath string) error {

	// Extract the file extension from the modelPath
	ext := strings.ToLower(filepath.Ext(modelPath))

	if len(*modelType) > 0 && slices.Contains(MODEL_TYPES, *modelType) {
		log.Warn(fmt.Sprintf("Model type %s not supported. Supported types are: %s", *modelType, MODEL_TYPES))
	}

	switch ext {
	case ".json":
		*modelType = string(tlApi.IMPORTMODELTYPE_JSON_TF2)
		return nil
	case ".h5":
		*modelType = string(tlApi.IMPORTMODELTYPE_H5_TF2)
		return nil
	case ".onnx":
		*modelType = string(tlApi.IMPORTMODELTYPE_ONNX)
		return nil
	case ".tar.gz":
		*modelType = string(tlApi.IMPORTMODELTYPE_PB_TF2)
		return nil
	}

	modelTypePrompt := &survey.Select{
		Message: "Select model type?",
		Options: MODEL_TYPES,
	}
	err := survey.AskOne(modelTypePrompt, modelType)
	if err != nil {
		return err
	}
	return nil
}
