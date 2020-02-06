package models

func (svc Service) GetMetaData() ResourceMeta {
	return ResourceMeta{
		Name:         svc.Name,
		TemplatePath: ServiceTemplatePath,
		Config: map[string]interface{}{
			"Name":      svc.Name,
			"Ports":     svc.Ports,
			"Type":      svc.Type,
			"Headless":  svc.Headless,
			"Labels":    svc.Labels,
			"Selectors": svc.Selectors,
		},
		Type: Services,
	}
}
