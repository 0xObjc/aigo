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
		return nil, nil
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

func CreateConfigFile(dir string) error {
	configFilePath := filepath.Join(dir, "aigo.yaml")
	if _, err := os.Stat(configFilePath); !os.IsNotExist(err) {
		return fmt.Errorf("config file already exists")
	}

	cfg := ConfigFile{
		Language:     "go",
		ExcludeFiles: []string{"AigoTemplate.md", "aigo.yaml"},
		IncludeFiles: []string{},
	}

	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	err = os.WriteFile(configFilePath, data, 0644)
	return err
}
