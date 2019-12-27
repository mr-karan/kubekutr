package models

func (ss StatefulSet) GetMetaData() ResourceMeta {
	return ResourceMeta{
		Name:         ss.Name,
		TemplatePath: StatefulsetTemplatePath,
		Config: map[string]interface{}{
			"Name":        ss.Name,
			"Labels":      ss.Labels,
			"ServiceName": ss.ServiceName,
			"Containers":  ss.Containers,
			"Volumes":     ss.Volumes,
		},
		ManifestPath: StatefulsetsDir,
	}
}
