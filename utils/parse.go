package utils

import (
	"fmt"
	"os"
	"text/template"

	"github.com/knadh/stuffbin"
)

// parse takes in a template path and the variables to be "applied" on it. The rendered template
// is saved to the destination path.
func parse(src string, dest string, config map[string]interface{}, fs stuffbin.FileSystem) error {
	// read template file
	tmpl := template.New(src)
	f, err := fs.Read(src)
	if err != nil {
		return fmt.Errorf("error reading template file %s: %v", src, err)
	}
	// convert file to a Template object
	_, err = tmpl.Parse(string(f))
	if err != nil {
		return err
	}
	// create an empty file
	output, err := os.Create(dest)
	defer output.Close()
	if err != nil {
		return err
	}
	// apply the variable and save the rendered template to the file.
	err = tmpl.Execute(output, config)
	if err != nil {
		return err
	}
	return nil
}

func saveResource(template string, name string, dest string, config map[string]interface{}, fs stuffbin.FileSystem) error {
	// parse template file and output yaml
	err := parse(template, dest, config, fs)
	if err != nil {
		return err
	}
	return nil
}
