package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/c-bata/go-prompt"
	"github.com/urfave/cli"
	"zerodha.tech/janus/models"
	"zerodha.tech/janus/utils"
)

var emptyComplete = func(prompt.Document) []prompt.Suggest { return []prompt.Suggest{} }

// ScaffoldProject creates an opinioated GitOps structure for Kubernetes manifests.
func (hub *Hub) ScaffoldProject(config models.Config) cli.Command {
	return cli.Command{
		Name:    "scaffold",
		Aliases: []string{"s"},
		Usage:   "Scaffold a new project with opinionated gitops structure",
		Action:  hub.initApp(hub.scaffold),
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "output, o",
				Usage: "Path to manifests output directory for `PROJECT`",
			},
		},
	}
}

func (hub *Hub) scaffold(cliCtx *cli.Context) error {
	var (
		parentDir = utils.GetRootDir(cliCtx.String("output"))
	)
	utils.CreateGitopsDirectory(subPaths, parentDir)
	// Create deployments
	for _, dep := range hub.Config.Deployments {
		err := loadDeployment(dep, filepath.Join(parentDir, BaseDir, DeploymentsDir))
		if err != nil {
			return err
		}
		hub.Logger.Debugf("Created manifest for deployment: %s", dep.Name)
	}
	// Create services
	for _, svc := range hub.Config.Services {
		err := loadService(svc, filepath.Join(parentDir, BaseDir, ServicesDir))
		if err != nil {
			return err
		}
		hub.Logger.Debugf("Created manifest for service: %s", svc.Name)
	}
	// Create ingress
	for _, ing := range hub.Config.Ingresses {
		err := loadIngress(ing, filepath.Join(parentDir, BaseDir, "ingresses"))
		if err != nil {
			return err
		}
		hub.Logger.Debugf("Created manifest for ingress: %s", ing.Name)
	}
	return nil
}

func loadDeployment(dep models.Deployment, dest string) error {
	// Deployment Config
	config := map[string]interface{}{
		"Name":       dep.Name,
		"Replicas":   dep.Replicas,
		"Containers": dep.Containers,
	}
	// read template file
	err := utils.Parse("templates/deployment.tmpl", filepath.Join(dest, fmt.Sprintf("%s.yml", dep.Name)), config)
	if err != nil {
		return err
	}
	return nil
}

func loadService(svc models.Service, dest string) error {
	// Service Config
	config := map[string]interface{}{
		"Name":       svc.Name,
		"Port":       svc.Port,
		"TargetPort": svc.TargetPort,
		"Type":       svc.Type,
	}
	// read template file
	err := utils.Parse("templates/service.tmpl", filepath.Join(dest, fmt.Sprintf("%s.yml", svc.Name)), config)
	if err != nil {
		return err
	}
	return nil
}

func loadIngress(ing models.Ingress, dest string) error {
	// Service Config
	config := map[string]interface{}{
		"Name":  ing.Name,
		"Class": ing.Class,
		"Paths": ing.Paths,
	}
	// read template file
	err := utils.Parse("templates/ingress.tmpl", filepath.Join(dest, fmt.Sprintf("%s.yml", ing.Name)), config)
	if err != nil {
		return err
	}
	return nil
}
