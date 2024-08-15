package config

import (
	"path/filepath"
	"strings"
)

type Config struct {
	Language     string
	ExcludeFiles []string
	IncludeFiles []string
}

func NewConfig() Config {
	return Config{
		Language:     "go",
		ExcludeFiles: []string{"aigo.yaml", "AigoTemplate.md"},
		IncludeFiles: []string{},
	}
}

func (c *Config) ShouldInclude(path string) bool {
	for _, pattern := range c.IncludeFiles {
		matched, err := filepath.Match(pattern, path)
		if err == nil && matched {
			return true
		}
	}
	return false
}

func (c *Config) ShouldExclude(path string) bool {
	// 先检查是否在 include_files 中
	if c.ShouldInclude(path) {
		return false
	}

	for _, pattern := range c.ExcludeFiles {
		matched, err := filepath.Match(pattern, filepath.Base(path))
		if err == nil && matched {
			return true
		}

		// 处理目录路径的匹配
		if strings.HasSuffix(pattern, "/") {
			matched, err = filepath.Match(pattern+"*", path)
			if err == nil && matched {
				return true
			}
		}
	}
	return false
}

func GetDefaultExcludeRules(language string) []string {
	return defaultExcludeRules[language]
}
