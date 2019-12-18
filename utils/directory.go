package utils

import (
	"os"
	"path/filepath"
)

// GetRootDir returns a relative path to the root of Gitops directory.
func GetRootDir(dest string) string {
	if dest == "" {
		dest, _ = os.Getwd()
	}
	return filepath.Join(dest)
}

// CreateGitopsDirectory ...
func CreateGitopsDirectory(subPaths []string, parentDir string) {
	for _, p := range subPaths {
		os.MkdirAll(filepath.Join(parentDir, "base", p), os.ModePerm)
	}
}

// LookupGitopsDirectory ...
func LookupGitopsDirectory(subPaths []string, parentDir string) error {
	for _, sub := range subPaths {
		if _, err := os.Stat(filepath.Join(parentDir, "base", sub)); os.IsNotExist(err) {
			return err
		}
	}
	return nil
}
