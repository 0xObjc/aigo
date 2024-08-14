package language

import (
	"os"
	"path/filepath"
)

type DefaultDetector struct{}

func (d *DefaultDetector) DetectLanguage(dir string) string {
	// 优先检测项目文件
	for file, language := range ProjectFiles {
		if _, err := os.Stat(filepath.Join(dir, file)); err == nil {
			return language
		}
	}

	// 其次检测文件扩展名
	var detectedLanguage string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			ext := filepath.Ext(path)
			if language, ok := LanguageExtensions[ext]; ok {
				detectedLanguage = language
				return filepath.SkipDir
			}
		}
		return nil
	})

	if err == nil && detectedLanguage != "" {
		return detectedLanguage
	}

	return "unknown"
}
