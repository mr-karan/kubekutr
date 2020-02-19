package cmd

import (
	"fmt"
	"os"

	"github.com/knadh/stuffbin"
	models "zerodha.tech/kubekutr/models"
	"zerodha.tech/kubekutr/utils"
)

func prepareResources(resources []models.Resource, projectDir string, workload string, fs stuffbin.FileSystem) error {
	for _, r := range resources {
		err := utils.CreateResource(r, projectDir, workload, fs)
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

// createDefaultConfig takes a default config template file and writes to the current directory
func createDefaultConfig(cfgFile []byte, fName string) error {
	f, err := os.Create(fName)
	if err != nil {
		return fmt.Errorf("error while creating sample config: %v", err)
	}
	_, err = f.Write(cfgFile)
	if err != nil {
		return fmt.Errorf("error while copying sample config: %v", err)
	}
	return nil
}
