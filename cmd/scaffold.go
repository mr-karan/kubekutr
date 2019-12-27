package cmd

import (
	"github.com/urfave/cli"
	"zerodha.tech/kubekutr/models"
	"zerodha.tech/kubekutr/utils"
)

// ScaffoldProject creates an opinioated GitOps structure for Kubernetes manifests.
func (hub *Hub) ScaffoldProject(config models.Config) cli.Command {
	return cli.Command{
		Name:    "scaffold",
		Aliases: []string{"s"},
		Usage:   "Scaffold a new project with gitops structure",
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
		projectDir = utils.GetRootDir(cliCtx.String("output"))
	)
	// Scaffold directory
	utils.CreateGitopsDirectory(subPaths, projectDir)
	// Create deployments
	resources := []models.Resource{}
	for _, dep := range hub.Config.Deployments {
		resources = append(resources, models.Resource(dep))
	}
	// Create services
	for _, svc := range hub.Config.Services {
		resources = append(resources, models.Resource(svc))
	}
	// Create ingress
	for _, ing := range hub.Config.Ingresses {
		resources = append(resources, models.Resource(ing))
	}
	// Create statefulset
	for _, ss := range hub.Config.StatefulSets {
		resources = append(resources, models.Resource(ss))
	}
	return prepareResources(resources, projectDir, hub.Fs)
}
