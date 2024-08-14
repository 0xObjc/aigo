package config

import (
	"path/filepath"
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

func (c *Config) ShouldExclude(path string) bool {
	for _, pattern := range c.ExcludeFiles {
		matched, err := filepath.Match(pattern, path)
		if err == nil && matched {
			return true
		}
	}
	return false
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

var defaultExcludeRules = map[string][]string{
	"go":     {"aigo.yaml", "AigoTemplate.md"},
	"python": {"__pycache__", "*.pyc"},
	// 添加其他语言的默认规则
}

func GetDefaultExcludeRules(language string) []string {
	return defaultExcludeRules[language]
}
