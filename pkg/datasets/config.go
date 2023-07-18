package datasets

import (
	"embed"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"text/template"
)

//go:embed template/*
var templateDir embed.FS
const configFileName = ".tensorleap.yaml"
type DatasetConfig struct {
	DatasetId       string   `yaml:"datasetId"`
	EntryFile       string   `yaml:"entryFile"`
	IncludePatterns []string `yaml:"include"`
}

type InitTemplateValues struct {
	DatasetId string
}

func CreateDatasetConfig(datasetId string) error {
	files, _ := templateDir.ReadDir("template")
	for _, f := range files {
		fileName := f.Name()
		if _, err := os.Stat(fileName); err == nil {
			fmt.Println(fileName, "already exists")
		} else {
			fmt.Println("Writing file:", fileName)
			templateContent, _ := templateDir.ReadFile("template/" + fileName)
			tmpl, _ := template.New(fileName).Parse(string(templateContent))
			targetFile, err := os.Create(fileName)
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

func GetDatasetConfig() (*DatasetConfig, error) {
	content, err := os.ReadFile(configFileName)
	if err != nil {
		return nil, err
	}
	datasetConfig := DatasetConfig{}
	err = yaml.Unmarshal(content, &datasetConfig)
	return &datasetConfig, err
}

func SetDatasetConfig(datasetConfig *DatasetConfig) error {
	content, err := yaml.Marshal(&datasetConfig)
	if err != nil {
		return err
	}
	err = os.WriteFile(configFileName, content, 0644)
	return err
}