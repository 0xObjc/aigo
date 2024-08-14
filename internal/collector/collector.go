package collector

import (
	"fmt"
	"github.com/0xObjc/aigo/internal/model"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/0xObjc/aigo/internal/config"
)

func CollectFiles(dir string, cfg config.Config) ([]model.FileWithLanguage, error) {
	var files []model.FileWithLanguage

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error walking directory: %v\n", err)
			return err
		}

		if info.IsDir() {
			if strings.HasPrefix(info.Name(), ".") {
				return filepath.SkipDir
			}
			return nil
		}

		relPath, err := filepath.Rel(dir, path)
		if err != nil {
			fmt.Printf("Error getting relative path: %v\n", err)
			return err
		}

		if cfg.ShouldExclude(relPath) {
			return nil
		}

		if cfg.ShouldInclude(relPath) || true {
			content, err := ioutil.ReadFile(path)
			if err != nil {
				fmt.Printf("Error reading file: %v\n", err)
				return err
			}

			language := getFileLanguage(relPath) // 获取文件语言

			files = append(files, model.FileWithLanguage{
				Name:     relPath,
				Content:  string(content),
				Language: language,
			})
		}

		return nil
	})

	return files, err
}

func getFileLanguage(path string) string {
	ext := filepath.Ext(path)
	switch ext {
	case ".go":
		return "go"
	case ".java":
		return "java"
	case ".py":
		return "python"
	// 添加其他语言的扩展名和对应的语言
	default:
		return ""
	}
}
