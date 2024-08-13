package renderer

import (
	"bytes"
	"github.com/0xObjc/aigo/internal/model"
	"io/ioutil"
	"text/template"

	"github.com/atotto/clipboard"
)

type TemplateData struct {
	ProjectStructure string
	Files            []model.File
}

func RenderTemplate(data TemplateData) error {
	templateBytes, err := ioutil.ReadFile("template.md")
	if err != nil {
		return err
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
