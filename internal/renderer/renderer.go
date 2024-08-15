package renderer

import (
	"bytes"
	"github.com/0xObjc/aigo/internal/model"
	"github.com/atotto/clipboard"
	"os"
	"path/filepath"
	"regexp"
	"text/template"
)

type TemplateData struct {
	ProjectStructure string
	Files            []model.FileWithLanguage // 修改为包含语言信息的文件结构
}

func RenderTemplate(dir string, data TemplateData) (string, int, error) {
	var templateBytes []byte
	var err error

	templatePath := filepath.Join(dir, "AigoTemplate.md")
	if _, err := os.Stat(templatePath); err == nil {
		templateBytes, err = os.ReadFile(templatePath)
		if err != nil {
			return "", 0, err
		}
	} else {
		templateBytes, err = defaultTemplate.ReadFile("AigoTemplate.md")
		if err != nil {
			return "", 0, err
		}
	}

	tmpl, err := template.New("project").Parse(string(templateBytes))
	if err != nil {
		return "", 0, err
	}

	var result bytes.Buffer
	err = tmpl.Execute(&result, data)
	if err != nil {
		return "", 0, err
	}

	renderedContent := result.String()
	tokenCount := estimateTokenCount(renderedContent)

	err = clipboard.WriteAll(renderedContent)
	if err != nil {
		return "", 0, err
	}

	return renderedContent, tokenCount, nil
}

func estimateTokenCount(content string) int {
	// 使用正则表达式匹配单词和标点符号
	re := regexp.MustCompile(`[\w']+|[.,!?;]`)
	tokens := re.FindAllString(content, -1)
	return len(tokens)
}
