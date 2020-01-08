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
	// Create deployments
	for _, workload := range hub.Config.Workloads {
		resources := []models.Resource{}
		for _, dep := range workload.Deployments {
			resources = append(resources, models.Resource(dep))
		}
		// Create services
		for _, svc := range workload.Services {
			resources = append(resources, models.Resource(svc))
		}
		// Create ingress
		for _, ing := range workload.Ingresses {
			resources = append(resources, models.Resource(ing))
		}
		// Create statefulset
		for _, ss := range workload.StatefulSets {
			resources = append(resources, models.Resource(ss))
		}
		// Scaffold directory
		utils.CreateGitopsDirectory(projectDir, workload.Name)
		err := prepareResources(resources, projectDir, workload.Name, hub.Fs)
		if err != nil {
			return err
		}
	}
	return nil
}
