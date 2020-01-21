package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/knadh/stuffbin"
	"zerodha.tech/kubekutr/models"
)

// CreateResource fetches metadata about the resource and produces the manifest.
func CreateResource(resource models.Resource, rootDir string, workload string, fs stuffbin.FileSystem) error {
	var (
		template           = resource.GetMetaData().TemplatePath
		name               = resource.GetMetaData().Name
		config             = resource.GetMetaData().Config
		fName              = fmt.Sprintf("%s-%s.yml", resource.GetMetaData().Name, resource.GetMetaData().Type)
		dest     io.Writer = os.Stdout
	)

	if rootDir != "" {
		path := filepath.Join(rootDir, models.BaseDir, workload, fName)
		f, err := os.Create(path)
		if err != nil {
			return err
		}

		dest = f
	}

	return saveResource(template, name, dest, config, fs)
}
