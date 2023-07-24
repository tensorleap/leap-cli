package code

import (
	"embed"
	"fmt"
	"os"
	"path"
	"text/template"

	"gopkg.in/yaml.v3"
)

//go:embed template/*
var templateDir embed.FS

const configFileName = ".tensorleap.yaml"

type DatasetConfig struct {
	DatasetId       string   `yaml:"datasetId"`
	EntryFile       string   `yaml:"entryFile"`
	IncludePatterns []string `yaml:"include"`
}

func NewDatasetConfig(datasetId string, entryFile string, files []string) *DatasetConfig {
	if len(entryFile) == 0 {
		entryFile = "tensorleap.py"
	}
	return &DatasetConfig{
		DatasetId:       datasetId,
		EntryFile:       entryFile,
		IncludePatterns: files,
	}
}

type InitTemplateValues struct {
	DatasetId string
}

func CreateCodeTemplate(datasetId string, outputDir string) error {
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
				DatasetId: datasetId,
			}); err != nil {
				return err
			}
		}
	}
	return nil
}

func GetCodeIntegrationConfig() (*DatasetConfig, error) {
	content, err := os.ReadFile(configFileName)
	if err != nil {
		return nil, err
	}
	datasetConfig := DatasetConfig{}
	err = yaml.Unmarshal(content, &datasetConfig)
	return &datasetConfig, err
}

func SetCodeIntegrationConfig(datasetConfig *DatasetConfig, outputDir string) error {
	content, err := yaml.Marshal(&datasetConfig)
	if err != nil {
		return err
	}
	fullPath := path.Join(outputDir, configFileName)
	err = os.WriteFile(fullPath, content, 0644)
	return err
}
