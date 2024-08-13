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

		if !info.IsDir() && (cfg.IncludeAllFiles || strings.HasSuffix(info.Name(), ".go") || info.Name() == "go.mod") {
			content, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			relPath, err := filepath.Rel(dir, path)
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
