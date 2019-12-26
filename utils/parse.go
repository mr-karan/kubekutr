package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/knadh/stuffbin"
)

func parse(src string, dest string, config map[string]interface{}, fs stuffbin.FileSystem) error {
	t, err := stuffbin.ParseTemplates(nil, fs, src)

	if err != nil {
		return err
	}

	f, err := os.Create(dest)
	defer f.Close()
	if err != nil {
		return err
	}

	err = t.Execute(f, config)
	if err != nil {
		return err
	}
	return nil
}

func saveResource(template string, name string, dest string, config map[string]interface{}, fs stuffbin.FileSystem) error {
	// parse template file and output yaml
	err := parse(template, filepath.Join(dest, fmt.Sprintf("%s.yml", name)), config, fs)
	if err != nil {
		return err
	}
	return nil
}
