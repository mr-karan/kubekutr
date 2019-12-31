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
			"Labels":     svc.Labels,
			"Headless":   svc.Headless,
			"Selectors":  svc.Selectors,
		},
		ManifestPath: ServicesDir,
	}
}
