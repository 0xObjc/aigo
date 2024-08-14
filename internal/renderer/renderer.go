package renderer

import (
	"bytes"
	"os"
	"path/filepath"
	"text/template"

	"github.com/0xObjc/aigo/internal/model"
	"github.com/atotto/clipboard"
)

type TemplateData struct {
	ProjectStructure string
	Files            []model.File
}

func RenderTemplate(dir string, data TemplateData) error {
	var templateBytes []byte
	var err error

	templatePath := filepath.Join(dir, "AigoTemplate.md")
	if _, err := os.Stat(templatePath); err == nil {
		templateBytes, err = os.ReadFile(templatePath)
		if err != nil {
			return err
		}
	} else {
		templateBytes, err = os.ReadFile("template.md")
		if err != nil {
			return err
		}
	}

	tmpl, err := template.New("project").Parse(string(templateBytes))
	if err != nil {
		return err
	}

	var result bytes.Buffer
	err = tmpl.Execute(&result, data)
	if err != nil {
		return err
	}

	err = clipboard.WriteAll(result.String())
	return err
}
