package utils

import (
	"fmt"
	"path/filepath"

	"github.com/knadh/stuffbin"
	"zerodha.tech/kubekutr/models"
)

// CreateResource fetches metadata about the resource and produces the manifest.
func CreateResource(resource models.Resource, rootDir string, workload string, fs stuffbin.FileSystem) error {
	var (
		template = resource.GetMetaData().TemplatePath
		name     = resource.GetMetaData().Name
		config   = resource.GetMetaData().Config
		fName    = fmt.Sprintf("%s-%s.yml", resource.GetMetaData().Name, resource.GetMetaData().Type)
		dest     = filepath.Join(rootDir, models.BaseDir, workload, fName)
	)
	return saveResource(template, name, dest, config, fs)
}
