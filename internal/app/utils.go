package app

import (
	"fmt"
	"os"
	"path/filepath"
)

func createDefaultTemplateFile(dir string) error {
	defaultTemplatePath := "template.md"
	targetTemplatePath := filepath.Join(dir, "AigoTemplate.md")

	if _, err := os.Stat(targetTemplatePath); err == nil {
		return fmt.Errorf("default template file already exists")
	}

	content, err := os.ReadFile(defaultTemplatePath)
	if err != nil {
		return err
	}

	err = os.WriteFile(targetTemplatePath, content, 0644)
	return err
}
