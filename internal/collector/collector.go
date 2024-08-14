package collector

import (
	"github.com/0xObjc/aigo/internal/model"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/0xObjc/aigo/internal/config"
)

func CollectFiles(dir string, cfg config.Config) ([]model.File, error) {
	var files []model.File

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && strings.HasPrefix(info.Name(), ".") {
			return filepath.SkipDir
		}

		relPath, err := filepath.Rel(dir, path)
		if err != nil {
			return err
		}

		if shouldExclude(relPath, cfg.ExcludeFiles) {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		if shouldInclude(relPath, cfg.IncludeFiles) || (cfg.IncludeAllFiles || strings.HasSuffix(info.Name(), ".go") || info.Name() == "go.mod") {
			content, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			files = append(files, model.File{
				Name:    relPath,
				Content: string(content),
			})
		}

		return nil
	})

	return files, err
}

func shouldExclude(path string, excludeFiles []string) bool {
	for _, exclude := range excludeFiles {
		if path == exclude {
			return true
		}
	}
	return false
}

func shouldInclude(path string, includeFiles []string) bool {
	for _, include := range includeFiles {
		if path == include {
			return true
		}
	}
	return false
}
