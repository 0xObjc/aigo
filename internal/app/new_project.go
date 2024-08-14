package app

import (
	"fmt"
	"os"

	"github.com/0xObjc/aigo/internal/config"
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

	err := config.CreateConfigFile(dir)
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
