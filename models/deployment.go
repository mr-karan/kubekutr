package models

func (dep Deployment) GetMetaData() ResourceMeta {
	return ResourceMeta{
		Name:         dep.Name,
		TemplatePath: DeploymentTemplatePath,
		Config: map[string]interface{}{
			"Name":       dep.Name,
			"Replicas":   dep.Replicas,
			"Containers": dep.Containers,
			"Labels":     dep.Labels,
		},
		ManifestPath: DeploymentsDir,
	}
}
