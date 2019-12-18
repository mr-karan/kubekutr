package cmd

import (
	models "zerodha.tech/janus/models"
	"zerodha.tech/janus/utils"
)

func prepareResources(resources []models.Resource, projectDir string) error {
	for _, r := range resources {
		// if name != "" {
		// 	if name != r.GetMetaData().Name {
		// 		continue
		// 	}
		// }
		err := utils.CreateResource(r, projectDir)
		if err != nil {
			return err
		}
	}
	return nil
}
