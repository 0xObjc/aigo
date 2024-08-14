package app

import (
	"fmt"
	"github.com/0xObjc/aigo/internal/config"
	lang "github.com/0xObjc/aigo/internal/language"
	"os"
)

func CreateNewProject(args []string) {
	var dir string
	if len(args) == 2 {
		dir, _ = os.Getwd()
	} else if len(args) == 3 {
		dir = args[2]
	} else {
		ShowHelp()
		return
	}

	var projectLanguage string
	for _, detector := range lang.Detectors() {
		projectLanguage = detector.DetectLanguage(dir)
		if projectLanguage != "unknown" {
			break
		}
	}

	if projectLanguage == "unknown" {
		fmt.Println("Unable to detect project language. Using default language: go")
		projectLanguage = "go"
	}

	err := config.CreateConfigFile(dir, projectLanguage)
	if err != nil {
		fmt.Println("Error creating config file:", err)
	} else {
		fmt.Println("Config file created successfully")
	}

	err = createDefaultTemplateFile(dir)
	if err != nil {
		fmt.Println("Error creating default template file:", err)
	} else {
		fmt.Println("Default template file created successfully")
	}
}
