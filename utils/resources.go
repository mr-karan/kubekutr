package utils

import (
	"path/filepath"

	"github.com/knadh/stuffbin"
	"zerodha.tech/kubekutr/models"
)

// CreateResource fetches metadata about the resource and produces the resultant manifest.
func CreateResource(resource models.Resource, rootDir string, fs stuffbin.FileSystem) error {
	var (
		template = resource.GetMetaData().TemplatePath
		name     = resource.GetMetaData().Name
		config   = resource.GetMetaData().Config
		dest     = filepath.Join(rootDir, models.BaseDir, resource.GetMetaData().ManifestPath)
	)
	return saveResource(template, name, dest, config, fs)
}
