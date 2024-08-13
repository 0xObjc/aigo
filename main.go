package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"github.com/atotto/clipboard"
)

type File struct {
	Name    string
	Content string
}

type TemplateData struct {
	ProjectStructure string
	Files            []File
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: GoProjectStructToClipboard <directory>")
		return
	}

	dir := os.Args[1]
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		fmt.Printf("Directory %s does not exist\n", dir)
		return
	}

	var projectStructure bytes.Buffer
	err := generateProjectStructure(dir, &projectStructure, "")
	if err != nil {
		fmt.Println("Error generating project structure:", err)
		return
	}

	files, err := collectGoFiles(dir)
	if err != nil {
		fmt.Println("Error collecting .go files:", err)
		return
	}

	data := TemplateData{
		ProjectStructure: projectStructure.String(),
		Files:            files,
	}

	templateBytes, err := ioutil.ReadFile("template.md")
	if err != nil {
		fmt.Println("Error reading template file:", err)
		return
	}

	tmpl, err := template.New("project").Parse(string(templateBytes))
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}

	var result bytes.Buffer
	err = tmpl.Execute(&result, data)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return
	}

	err = clipboard.WriteAll(result.String())
	if err != nil {
		fmt.Println("Error writing to clipboard:", err)
		return
	}

	fmt.Println("Project structure and .go files copied to clipboard")
}

func generateProjectStructure(dir string, buf *bytes.Buffer, prefix string) error {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	// Sort entries to ensure directories come first
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].IsDir() && !entries[j].IsDir()
	})

	for i, entry := range entries {
		if strings.HasPrefix(entry.Name(), ".") {
			continue
		}

		if i == len(entries)-1 {
			buf.WriteString(prefix + "└── " + entry.Name())
		} else {
			buf.WriteString(prefix + "├── " + entry.Name())
		}

		if entry.IsDir() {
			buf.WriteString("/\n")
			newPrefix := prefix + "│   "
			if i == len(entries)-1 {
				newPrefix = prefix + "    "
			}
			err := generateProjectStructure(filepath.Join(dir, entry.Name()), buf, newPrefix)
			if err != nil {
				return err
			}
		} else if strings.HasSuffix(entry.Name(), ".go") {
			buf.WriteString("\n")
		} else {
			buf.WriteString("\n")
		}
	}

	return nil
}

func collectGoFiles(dir string) ([]File, error) {
	var files []File

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
			content, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			relPath, err := filepath.Rel(dir, path)
			if err != nil {
				return err
			}

			files = append(files, File{
				Name:    relPath,
				Content: string(content),
			})
		}

		return nil
	})

	return files, err
}
