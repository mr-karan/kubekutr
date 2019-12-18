package utils

import (
	"path/filepath"

	"zerodha.tech/janus/models"
)

// CreateResource ...
func CreateResource(resource models.Resource, rootDir string) error {
	var (
		template = resource.GetMetaData().TemplatePath
		name     = resource.GetMetaData().Name
		config   = resource.GetMetaData().Config
		dest     = filepath.Join(rootDir, models.BaseDir, resource.GetMetaData().ManifestPath)
	)
	return saveResource(template, name, dest, config)
}
