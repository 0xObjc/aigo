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
		fmt.Println(err)
		return
	}

	if configFile != nil {
		cfg.Language = configFile.Language
		cfg.ExcludeFiles = configFile.ExcludeFiles
		cfg.IncludeFiles = configFile.IncludeFiles
	}

	projectStructure, err := generator.GenerateProjectStructure(dir, cfg)
	if err != nil {
		fmt.Println("错误：生成项目结构失败:", err)
		return
	}

	files, err := collector.CollectFiles(dir, cfg)
	if err != nil {
		fmt.Println("错误：收集文件失败:", err)
		return
	}

	data := renderer.TemplateData{
		ProjectStructure: projectStructure,
		Files:            files,
	}

	_, tokenCount, err := renderer.RenderTemplate(dir, data)
	if err != nil {
		fmt.Println("错误：渲染模板失败:", err)
		return
	}

	// 仅打印项目结构
	fmt.Println(projectStructure)
	fmt.Printf("估计的 token 数量: %d\n", tokenCount)
	fmt.Println("模板结构已复制到剪贴板")
}
