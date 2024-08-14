package renderer

import (
	"bytes"
	"github.com/0xObjc/aigo/internal/model"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"

	"github.com/atotto/clipboard"
)

type TemplateData struct {
	ProjectStructure string
	Files            []model.File
}

func RenderTemplate(dir string, data TemplateData) error {
	var templateBytes []byte
	var err error

	// 检查是否存在自定义模板文件
	templatePath := filepath.Join(dir, "AigoTemplate.md")
	if _, err := os.Stat(templatePath); err == nil {
		templateBytes, err = ioutil.ReadFile(templatePath)
		if err != nil {
			return err
		}
	} else {
		// 使用默认模板文件
		templateBytes, err = ioutil.ReadFile("template.md")
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
