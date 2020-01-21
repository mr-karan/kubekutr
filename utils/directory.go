package utils

import (
	"os"
	"path/filepath"
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

// LookupGitopsDirectory checks if a directory with the same path already exists or not.
func LookupGitopsDirectory(subPaths []string, parentDir string) error {
	for _, sub := range subPaths {
		if _, err := os.Stat(filepath.Join(parentDir, "base", sub)); os.IsNotExist(err) {
			return err
		}
	}
	return nil
}
