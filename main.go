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

type Config struct {
	IncludeAllFiles bool
}

func main() {
	if len(os.Args) != 2 && len(os.Args) != 3 {
		fmt.Println("Usage: GoProjectStructToClipboard <directory> [-all]")
		return
	}

	dir := os.Args[1]
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		fmt.Printf("Directory %s does not exist\n", dir)
		return
	}

	config := Config{
		IncludeAllFiles: len(os.Args) == 3 && os.Args[2] == "-all",
	}

	// Print all arguments to terminal
	fmt.Println("Arguments:")
	for i, arg := range os.Args {
		fmt.Printf("Arg %d: %s\n", i, arg)
	}

	projectStructure, err := GenerateProjectStructure(dir, config)
	if err != nil {
		fmt.Println("Error generating project structure:", err)
		return
	}

	files, err := CollectGoFiles(dir)
	if err != nil {
		fmt.Println("Error collecting .go files:", err)
		return
	}

	data := TemplateData{
		ProjectStructure: projectStructure,
		Files:            files,
	}

	// Print project structure to terminal
	fmt.Println("Project Structure:")
	fmt.Println(projectStructure)

	// Render template and copy to clipboard
	err = RenderTemplate(data)
	if err != nil {
		fmt.Println("Error rendering template:", err)
		return
	}

	fmt.Println("Template structure copied to clipboard")
}

func GenerateProjectStructure(dir string, config Config) (string, error) {
	var buf bytes.Buffer
	err := generateProjectStructure(dir, &buf, "", config)
	return buf.String(), err
}

func generateProjectStructure(dir string, buf *bytes.Buffer, prefix string, config Config) error {
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
			err := generateProjectStructure(filepath.Join(dir, entry.Name()), buf, newPrefix, config)
			if err != nil {
				return err
			}
		} else if config.IncludeAllFiles || strings.HasSuffix(entry.Name(), ".go") {
			buf.WriteString("\n")
		} else {
			buf.WriteString("\n")
		}
	}

	return nil
}

func CollectGoFiles(dir string) ([]File, error) {
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
