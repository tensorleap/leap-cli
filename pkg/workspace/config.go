package workspace

import (
	"embed"
	"fmt"
	"os"
	"path"
	"text/template"

	"gopkg.in/yaml.v3"
	"k8s.io/kubectl/pkg/util/slice"
)

//go:embed template/*
var templateDir embed.FS

const configFileName = "leap.yaml"

type WorkspaceConfig struct {
	CodeIntegrationId string `yaml:"codeIntegrationId"`
	ProjectId         string `yaml:"projectId"`
	SecretId          string `yaml:"secretId"`
	// deprecated
	SecretManagerId string   `yaml:"secretManagerId,omitempty"`
	EntryFile       string   `yaml:"entryFile"`
	IncludePatterns []string `yaml:"include"`
}

func NewWorkspaceConfig(codeIntegrationId, projectId, entryFile, secretId string, files []string) *WorkspaceConfig {
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
		IncludePatterns:   files,
	}
}

type InitTemplateValues struct {
	CodeIntegrationId string
	ProjectId         string
	SecretId          string
}

func CreateCodeTemplate(codeIntegrationId, projectId, secretId, outputDir string) error {
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
			filePath := path.Join("template", fileName)
			templateContent, _ := templateDir.ReadFile(filePath)
			tmpl, _ := template.New(fileName).Parse(string(templateContent))
			targetPath := path.Join(outputDir, fileName)
			targetFile, err := os.Create(targetPath)
			if err != nil {
				return err
			}
			defer targetFile.Close()

			if err := tmpl.Execute(targetFile, &InitTemplateValues{
				CodeIntegrationId: codeIntegrationId,
				ProjectId:         projectId,
				SecretId:          secretId,
			}); err != nil {
				return err
			}
		}
	}
	return nil
}

func OverrideWorkspaceConfig(codeIntegrationId, projectId, entryFile, secretId, outputDir string) error {
	wc, err := GetWorkspaceConfig()
	if err != nil {
		wc = NewWorkspaceConfig(codeIntegrationId, projectId, entryFile, secretId, nil)
	} else {
		wc.CodeIntegrationId = codeIntegrationId
		wc.ProjectId = projectId
		wc.EntryFile = entryFile
		wc.SecretId = secretId
	}
	return SetWorkspaceConfig(wc, outputDir)
}

func GetWorkspaceConfig() (*WorkspaceConfig, error) {
	content, err := os.ReadFile(configFileName)
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
	fullPath := path.Join(outputDir, configFileName)
	err = os.WriteFile(fullPath, content, 0644)
	if err != nil {
		return fmt.Errorf("failed to update config (.leap.yaml) : %v", err)
	}
	return nil
}

func IsWorkspaceDir() bool {
	_, err := os.Stat(configFileName)
	return err == nil
}
