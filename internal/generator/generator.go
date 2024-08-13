package generator

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/0xObjc/aigo/internal/config"
)

func GenerateProjectStructure(dir string, cfg config.Config) (string, error) {
	var buf bytes.Buffer

	// Read go.mod to get project name
	goModPath := filepath.Join(dir, "go.mod")
	content, err := ioutil.ReadFile(goModPath)
	if err != nil {
		return "", err
	}
	projectName := extractProjectName(string(content))

	buf.WriteString(projectName + "\n")
	err = generateProjectStructure(dir, &buf, "", cfg)
	return buf.String(), err
}

func extractProjectName(goModContent string) string {
	lines := strings.Split(goModContent, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "module "))
		}
	}
	return "UnknownProject"
}

func generateProjectStructure(dir string, buf *bytes.Buffer, prefix string, cfg config.Config) error {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	// Sort entries to ensure directories come first
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].IsDir() && !entries[j].IsDir()
	})

	var filteredEntries []os.FileInfo
	for _, entry := range entries {
		if strings.HasPrefix(entry.Name(), ".") {
			continue
		}
		if cfg.IncludeAllFiles || entry.IsDir() || strings.HasSuffix(entry.Name(), ".go") || strings.HasSuffix(entry.Name(), "go.mod") {
			filteredEntries = append(filteredEntries, entry)
		}
	}

	for i, entry := range filteredEntries {
		if i == len(filteredEntries)-1 {
			buf.WriteString(prefix + "└── " + entry.Name())
		} else {
			buf.WriteString(prefix + "├── " + entry.Name())
		}

		if entry.IsDir() {
			buf.WriteString("/\n")
			newPrefix := prefix + "│   "
			if i == len(filteredEntries)-1 {
				newPrefix = prefix + "    "
			}
			err := generateProjectStructure(filepath.Join(dir, entry.Name()), buf, newPrefix, cfg)
			if err != nil {
				return err
			}
		} else {
			buf.WriteString("\n")
		}
	}

	return nil
}
