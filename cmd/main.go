package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/0xObjc/aigo/internal/collector"
	"github.com/0xObjc/aigo/internal/config"
	"github.com/0xObjc/aigo/internal/generator"
	"github.com/0xObjc/aigo/internal/renderer"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: aigo <directory> [-all]")
		return
	}

	command := os.Args[1]
	switch command {
	case "new":
		var dir string
		if len(os.Args) == 2 {
			// 如果没有提供目录参数，使用当前工作目录
			dir, _ = os.Getwd()
		} else if len(os.Args) == 3 {
			dir = os.Args[2]
		} else {
			fmt.Println("Usage: aigo new [directory]")
			return
		}

		err := config.CreateConfigFile(dir)
		if err != nil {
			fmt.Println("Error creating config file:", err)
		} else {
			fmt.Println("Config file created successfully")
		}

		// 创建默认模板文件
		err = createDefaultTemplateFile(dir)
		if err != nil {
			fmt.Println("Error creating default template file:", err)
		} else {
			fmt.Println("Default template file created successfully")
		}
		return
	default:
		dir := command
		cfg := config.NewConfig(os.Args)

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
}

func createDefaultTemplateFile(dir string) error {
	defaultTemplatePath := "template.md"
	targetTemplatePath := filepath.Join(dir, "AigoTemplate.md")

	// 检查目标目录中是否已存在默认模板文件
	if _, err := os.Stat(targetTemplatePath); err == nil {
		return fmt.Errorf("default template file already exists")
	}

	// 读取默认模板文件内容
	content, err := ioutil.ReadFile(defaultTemplatePath)
	if err != nil {
		return err
	}

	// 将默认模板文件内容写入目标目录
	err = ioutil.WriteFile(targetTemplatePath, content, 0644)
	return err
}
