package utils

import (
	"os"
	"text/template"
)

func Parse(src string, dest string, config map[string]interface{}) error {
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
