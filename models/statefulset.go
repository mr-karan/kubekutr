package models

func (ss StatefulSet) GetMetaData() ResourceMeta {
	return ResourceMeta{
		Name:         ss.Name,
		TemplatePath: StatefulsetTemplatePath,
		Config: map[string]interface{}{
			"Name":       ss.Name,
			// "Replicas":   ss.Replicas,
			// "Containers": ss.Containers,
			// "Labels":     ss.Labels,
		},
		ManifestPath: StatefulsetsDir,
	}
}
