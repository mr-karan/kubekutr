package cmd

import (
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
	for _, dep := range hub.Config.Deployments {
		err := utils.CreateResource(dep, parentDir)
		if err != nil {
			return err
		}
		hub.Logger.Debugf("Created manifest for deployment: %s", dep.Name)
	}

	// Create services
	for _, svc := range hub.Config.Services {
		err := utils.CreateResource(svc, parentDir)
		if err != nil {
			return err
		}
		hub.Logger.Debugf("Created manifest for service: %s", svc.Name)
	}
	// Create ingress
	for _, ing := range hub.Config.Ingresses {
		err := utils.CreateResource(ing, parentDir)
		if err != nil {
			return err
		}
		hub.Logger.Debugf("Created manifest for ingress: %s", ing.Name)
	}
	return nil
}
