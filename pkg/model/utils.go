package model

import (
	"path/filepath"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/tensorleap/leap-cli/pkg/log"
	"k8s.io/utils/strings/slices"
)

var MODEL_TYPES = []string{"JSON_TF2", "ONNX", "PB_TF2", "H5_TF2"}

func SelectModelType(modelType *string, modelPath string) {

	// Extract the file extension from the modelPath
	ext := strings.ToLower(filepath.Ext(modelPath))

	switch ext {
	case ".json":
		*modelType = "JSON_TF2"
		return
	case ".h5":
		*modelType = "H5_TF2"
		return
	default:
		if slices.Contains(MODEL_TYPES, *modelType) {
			return
		}
		log.Warn("Invalid model type")
	}

	modelTypePrompt := &survey.Select{
		Message: "Select model type?",
		Options: MODEL_TYPES,
	}
	_ = survey.AskOne(modelTypePrompt, modelType)
}
