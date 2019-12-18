package models

func (svc Service) GetMetaData() ResourceMeta {
	return ResourceMeta{
		Name:         svc.Name,
		TemplatePath: ServiceTemplatePath,
		Config: map[string]interface{}{
			"Name":       svc.Name,
			"Port":       svc.Port,
			"TargetPort": svc.TargetPort,
			"Type":       svc.Type,
		},
		ManifestPath: ServicesDir,
	}
}
