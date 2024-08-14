package app

import (
	"fmt"
)

func ShowHelp() {
	fmt.Println("用法: aigo <命令> [选项]")
	fmt.Println("命令:")
	fmt.Println("  new [目录]   创建新的配置文件和默认模板文件")
	fmt.Println("  w [目录]     生成指定目录的项目结构并复制到剪贴板")
	fmt.Println("  help         显示帮助信息")
}
