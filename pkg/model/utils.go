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
