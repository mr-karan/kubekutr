package cmd

import (
	"github.com/knadh/stuffbin"
	models "zerodha.tech/kubekutr/models"
	"zerodha.tech/kubekutr/utils"
)

func prepareResources(resources []models.Resource, projectDir string, fs stuffbin.FileSystem) error {
	for _, r := range resources {
		err := utils.CreateResource(r, projectDir, fs)
		if err != nil {
			return err
		}
	}
	return nil
}

func matchResource(input string, existing string) bool {
	if input == existing {
		return true
	}
	return false
}
