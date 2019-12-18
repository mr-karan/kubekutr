package models

func (ing Ingress) GetMetaData() ResourceMeta {
	return ResourceMeta{
		Name:         ing.Name,
		TemplatePath: IngressTemplatePath,
		Config: map[string]interface{}{
			"Name":  ing.Name,
			"Class": ing.Class,
			"Paths": ing.Paths,
		},
		ManifestPath: IngressesDir,
	}
}
