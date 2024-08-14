package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type ConfigFile struct {
	Language     string   `yaml:"language"`
	ExcludeFiles []string `yaml:"exclude_files"`
	IncludeFiles []string `yaml:"include_files"`
}

func LoadConfigFile(dir string) (*ConfigFile, error) {
	configFilePath := filepath.Join(dir, "aigo.yaml")
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		// 如果没有找到配置文件，返回一个友好的提示
		return nil, fmt.Errorf("未找到配置文件。请使用 'aigo new' 创建一个。")
	}

	content, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}

	var cfg ConfigFile
	err = yaml.Unmarshal(content, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func CreateConfigFile(dir string, language string) error {
	configFilePath := filepath.Join(dir, "aigo.yaml")
	if _, err := os.Stat(configFilePath); !os.IsNotExist(err) {
		return fmt.Errorf("配置文件已存在")
	}

	excludeFiles := append(GetDefaultExcludeRules(language), "aigo.yaml", "AigoTemplate.md")

	cfg := ConfigFile{
		Language:     language,
		ExcludeFiles: excludeFiles,
		IncludeFiles: []string{},
	}

	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	err = os.WriteFile(configFilePath, data, 0644)
	return err
}
