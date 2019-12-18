package models

import "fmt"

const (
	BaseDir         = "base"
	DeploymentsDir  = "deployments"
	ServicesDir     = "services"
	IngressesDir    = "ingresses"
	StatefulsetsDir = "statefulsets"
	RBACDir         = "rbac"
	TemplatesDir    = "templates"
)

var (
	IngressTemplatePath     = fmt.Sprintf("%s/ingress.tmpl", TemplatesDir)
	DeploymentTemplatePath  = fmt.Sprintf("%s/deployment.tmpl", TemplatesDir)
	ServiceTemplatePath     = fmt.Sprintf("%s/service.tmpl", TemplatesDir)
	StatefulsetTemplatePath = fmt.Sprintf("%s/statefulset.tmpl", TemplatesDir)
)
