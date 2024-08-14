package collector

import (
	"fmt"
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
			fmt.Printf("Error walking directory: %v\n", err)
			return err
		}

		if info.IsDir() {
			if strings.HasPrefix(info.Name(), ".") {
				return filepath.SkipDir
			}
			return nil
		}

		relPath, err := filepath.Rel(dir, path)
		if err != nil {
			fmt.Printf("Error getting relative path: %v\n", err)
			return err
		}

		if cfg.ShouldExclude(relPath) {
			return nil
		}

		if cfg.ShouldInclude(relPath) || true {
			content, err := ioutil.ReadFile(path)
			if err != nil {
				fmt.Printf("Error reading file: %v\n", err)
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
