package app

import (
	"fmt"
)

func ShowHelp() {
	fmt.Println("Usage: aigo <command> [options]")
	fmt.Println("Commands:")
	fmt.Println("  new [directory]   创建新的配置文件和默认模板文件")
	fmt.Println("  w [directory]     生成指定目录的项目结构并复制到剪贴板")
	fmt.Println("  help              显示帮助信息")
}
