package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

func parse(src string, dest string, config map[string]interface{}) error {
	t, err := template.ParseFiles(src)
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

func saveResource(template string, name string, dest string, config map[string]interface{}) error {
	// parse template file and output yaml
	err := parse(template, filepath.Join(dest, fmt.Sprintf("%s.yml", name)), config)
	if err != nil {
		return err
	}
	return nil
}
