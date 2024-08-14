package generator

import (
	"bytes"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/0xObjc/aigo/internal/config"
)

func GenerateProjectStructure(dir string, cfg config.Config) (string, error) {
	var buf bytes.Buffer

	goModPath := filepath.Join(dir, "go.mod")
	content, err := os.ReadFile(goModPath)
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
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].IsDir() && !entries[j].IsDir()
	})

	var filteredEntries []os.DirEntry
	for _, entry := range entries {
		if strings.HasPrefix(entry.Name(), ".") {
			continue
		}
		if cfg.ShouldInclude(entry.Name()) || true {
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
