package main

import (
	"fmt"
	"os"

	"github.com/0xObjc/aigo/internal/collector"
	"github.com/0xObjc/aigo/internal/config"
	"github.com/0xObjc/aigo/internal/generator"
	"github.com/0xObjc/aigo/internal/renderer"
)

func main() {
	if len(os.Args) != 2 && len(os.Args) != 3 {
		fmt.Println("Usage: aigo <directory> [-all]")
		return
	}

	dir := os.Args[1]
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		fmt.Printf("Directory %s does not exist\n", dir)
		return
	}

	cfg := config.NewConfig(os.Args)

	// Print all arguments to terminal
	fmt.Println("Arguments:")
	for i, arg := range os.Args {
		fmt.Printf("Arg %d: %s\n", i, arg)
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

	// Print project structure to terminal
	fmt.Println("Project Structure:")
	fmt.Println(projectStructure)

	// Render template and copy to clipboard
	err = renderer.RenderTemplate(data)
	if err != nil {
		fmt.Println("Error rendering template:", err)
		return
	}

	fmt.Println("Template structure copied to clipboard")
}
