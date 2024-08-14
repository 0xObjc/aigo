package app

import (
	"fmt"
	"os"

	"github.com/0xObjc/aigo/internal/collector"
	"github.com/0xObjc/aigo/internal/config"
	"github.com/0xObjc/aigo/internal/generator"
	"github.com/0xObjc/aigo/internal/renderer"
)

func GenerateProjectStructure(args []string) {
	var dir string
	if len(args) == 2 {
		dir, _ = os.Getwd()
	} else if len(args) == 3 {
		dir = args[2]
	} else {
		ShowHelp()
		return
	}

	cfg := config.NewConfig()

	configFile, err := config.LoadConfigFile(dir)
	if err != nil {
		fmt.Println("Error loading config file:", err)
		return
	}

	if configFile != nil {
		cfg.Language = configFile.Language
		cfg.ExcludeFiles = configFile.ExcludeFiles
		cfg.IncludeFiles = configFile.IncludeFiles
	}

	projectStructure, err := generator.GenerateProjectStructure(dir, cfg)
	if err != nil {
		fmt.Println("Error generating project structure:", err)
		return
	}

	files, err := collector.CollectFiles(dir, cfg)
	if err != nil {
		fmt.Println("Error collecting files:", err)
		return
	}

	data := renderer.TemplateData{
		ProjectStructure: projectStructure,
		Files:            files,
	}

	err = renderer.RenderTemplate(dir, data)
	if err != nil {
		fmt.Println("Error rendering template:", err)
		return
	}

	fmt.Println("Template structure copied to clipboard")
}
