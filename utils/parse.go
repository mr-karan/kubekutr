package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/knadh/stuffbin"
)

// parse takes in a template path and the variables to be "applied" on it. The rendered template
// is saved to the destination path.
func parse(src string, dest string, config map[string]interface{}, fs stuffbin.FileSystem) error {
	// load template
	tmpl, err := stuffbin.ParseTemplates(nil, fs, src)
	if err != nil {
		return err
	}
	// create an empty file
	f, err := os.Create(dest)
	defer f.Close()
	if err != nil {
		return err
	}
	// apply the variable and save the rendered template to the file.
	err = tmpl.Execute(f, config)
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
