package config

import (
	"embed"
	"fmt"
	"os"
	"text/template"

	"github.com/spf13/cobra"
)

//go:embed template/*
var templateDir embed.FS

type InitTemplateValues struct {
	DatasetId string
}

func CreateDatasetConfig(datasetId string) {
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
			cobra.CheckErr(err)

			err = tmpl.Execute(targetFile, &InitTemplateValues{
				DatasetId: datasetId,
			})
			cobra.CheckErr(err)
		}
	}
}
