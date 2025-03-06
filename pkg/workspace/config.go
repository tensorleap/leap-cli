package workspace

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"gopkg.in/yaml.v3"
	"k8s.io/kubectl/pkg/util/slice"
)

//go:embed template/*
var templateDir embed.FS

const CONFIG_FILE_NAME = "leap.yaml"

const UNDEFINED_WORKSPACE_CONFIG = "undefined"

type WorkspaceConfig struct {
	CodeIntegrationId string `yaml:"codeIntegrationId"`
	ProjectId         string `yaml:"projectId"`
	SecretId          string `yaml:"secretId"`
	Branch            string `yaml:"branch,omitempty"`
	// deprecated
	SecretManagerId string   `yaml:"secretManagerId,omitempty"`
	EntryFile       string   `yaml:"entryFile"`
	IncludePatterns []string `yaml:"include"`
	PythonVersion   string   `yaml:"pythonVersion,omitempty"`
}

func NewWorkspaceConfig(codeIntegrationId, projectId, entryFile, secretId, branch, pythonVersion string, files []string) *WorkspaceConfig {
	if len(entryFile) == 0 {
		entryFile = "leap_binder.py"
	}
	if !slice.ContainsString(files, entryFile, nil) {
		files = append(files, entryFile)
	}
	return &WorkspaceConfig{
		ProjectId:         projectId,
		CodeIntegrationId: codeIntegrationId,
		EntryFile:         entryFile,
		SecretId:          secretId,
		Branch:            branch,
		PythonVersion:     pythonVersion,
		IncludePatterns:   files,
	}
}

type InitTemplateValues struct {
	CodeIntegrationId string
	ProjectId         string
	SecretId          string
	Branch            string
	PythonVersion     string
}

func CreateCodeTemplate(codeIntegrationId, projectId, secretId, branch, outputDir, pythonVersionId string) error {
	// Create the directory for the file if it doesn't exist
	err := os.MkdirAll(outputDir, 0755)
	if err != nil {
		return err
	}

	files, _ := templateDir.ReadDir("template")
	for _, f := range files {
		fileName := f.Name()
		if _, err := os.Stat(fileName); err == nil {
			fmt.Println(fileName, "already exists")
		} else {
			fmt.Println("Adding file:", fileName)
			filePath := filepath.Join("template", fileName)
			templateContent, _ := templateDir.ReadFile(filePath)
			tmpl, _ := template.New(fileName).Parse(string(templateContent))
			targetPath := filepath.Join(outputDir, fileName)
			targetFile, err := os.Create(targetPath)
			if err != nil {
				return err
			}
			defer targetFile.Close()

			if err := tmpl.Execute(targetFile, &InitTemplateValues{
				CodeIntegrationId: codeIntegrationId,
				ProjectId:         projectId,
				SecretId:          secretId,
				Branch:            branch,
				PythonVersion:     pythonVersionId,
			}); err != nil {
				return err
			}
		}
	}
	return nil
}

func OverrideWorkspaceConfig(codeIntegrationId, projectId, entryFile, secretId, branch string, files []string, outputDir string) error {
	wc, err := GetWorkspaceConfig()
	if err != nil {
		wc = NewWorkspaceConfig(codeIntegrationId, projectId, entryFile, secretId, branch, files)
	} else {
		wc.CodeIntegrationId = codeIntegrationId
		wc.ProjectId = projectId
		wc.EntryFile = entryFile
		wc.SecretId = secretId
	}
	return SetWorkspaceConfig(wc, outputDir)
}

func GetWorkspaceConfig() (*WorkspaceConfig, error) {
	content, err := os.ReadFile(CONFIG_FILE_NAME)
	if os.IsNotExist(err) {
		return nil, err
	}

	if err != nil {
		return nil, fmt.Errorf("not found config file, please make sure to init: %v", err)
	}
	workspaceConfig := WorkspaceConfig{}
	err = yaml.Unmarshal(content, &workspaceConfig)
	// handle deprecated SecretManagerId
	if len(workspaceConfig.SecretId) == 0 {
		workspaceConfig.SecretId = workspaceConfig.SecretManagerId
	}
	if len(workspaceConfig.SecretManagerId) > 0 {
		workspaceConfig.SecretManagerId = ""
		_ = SetWorkspaceConfig(&workspaceConfig, ".")
	}
	return &workspaceConfig, err
}

func SetWorkspaceConfig(workspaceConfig *WorkspaceConfig, outputDir string) error {
	content, err := yaml.Marshal(&workspaceConfig)
	if err != nil {
		return err
	}
	fullPath := filepath.Join(outputDir, CONFIG_FILE_NAME)
	err = os.WriteFile(fullPath, content, 0644)
	if err != nil {
		return fmt.Errorf("failed to update config (%s) : %v", CONFIG_FILE_NAME, err)
	}
	return nil
}

func IsWorkspaceDir() bool {
	_, err := os.Stat(CONFIG_FILE_NAME)
	return err == nil
}
