package models

import "fmt"

const (
	BaseDir      = "base"
	Deployments  = "deployment"
	Services     = "service"
	Ingresses    = "ingress"
	Statefulsets = "statefulset"
	TemplatesDir = "templates"
)

var (
	IngressTemplatePath     = fmt.Sprintf("%s/ingress.tmpl", TemplatesDir)
	DeploymentTemplatePath  = fmt.Sprintf("%s/deployment.tmpl", TemplatesDir)
	ServiceTemplatePath     = fmt.Sprintf("%s/service.tmpl", TemplatesDir)
	StatefulsetTemplatePath = fmt.Sprintf("%s/statefulset.tmpl", TemplatesDir)
	KustomizeTemplatePath   = fmt.Sprintf("%s/kustomization.tmpl", TemplatesDir)
)
