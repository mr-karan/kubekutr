package utils

import (
	"os"
	"path/filepath"

	"github.com/knadh/stuffbin"
	"github.com/mr-karan/kubekutr/models"
)

// GetRootDir returns a relative path to the root of Gitops directory.
func GetRootDir(dest string) string {
	if dest == "" {
		return ""
	}
	return filepath.Clean(dest)
}

// CreateGitopsDirectory creates an opinionated directory structure to organize
// resource manifests efficiently. The directory is ideally to be used with Kustomize
// as a "base".
func CreateGitopsDirectory(parentDir string, workload string) {
	os.MkdirAll(filepath.Join(parentDir, "base", workload), os.ModePerm)
}

// CreateKustomization does something
func CreateKustomization(parentDir string, rNames []string, fs stuffbin.FileSystem) error {
	path := filepath.Join(parentDir, "kustomization.yml")
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	var cfg = map[string]interface{}{
		"Paths": rNames,
	}
	return saveResource(models.KustomizeTemplatePath, "kustomization.yml", f, cfg, fs)
}

// LookupGitopsDirectory checks if a directory with the same path already exists or not.
func LookupGitopsDirectory(subPaths []string, parentDir string) error {
	for _, sub := range subPaths {
		if _, err := os.Stat(filepath.Join(parentDir, "base", sub)); os.IsNotExist(err) {
			return err
		}
	}
	return nil
}
