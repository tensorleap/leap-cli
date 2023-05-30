package config

import (
	"embed"
	"fmt"
	"os"
)

//go:embed template/*
var templateDir embed.FS

func CreateDatasetConfig() {
	files, _ := templateDir.ReadDir("template")
	for _, f := range files {
		fileName := f.Name()
		if _, err := os.Stat(fileName); err == nil {
			fmt.Println(fileName, "already exists")
		} else {
			fmt.Println("Writing file:", fileName)
			content, _ := templateDir.ReadFile("template/" + fileName)
			os.WriteFile(fileName, content, 0666)
		}
	}
}
