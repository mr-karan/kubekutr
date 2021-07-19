package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/go-yaml/yaml"
	"github.com/mr-karan/kubekutr/models"
	"github.com/urfave/cli"
)

var (
	configName = "kubekutr.yml"
)

// InitProject initializes git repo and copies a sample config
func (hub *Hub) InitProject(config models.Config) cli.Command {
	return cli.Command{
		Name:    "init",
		Aliases: []string{"i"},
		Usage:   "Initializes an empty git repo with a kubekutr config file.",
		Action:  hub.init,
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "default, d",
				Usage: "Use the default config file",
			},
			cli.StringFlag{
				Name:  "output, o",
				Usage: "Config file name.",
			},
		},
	}
}

func (hub *Hub) init(cliCtx *cli.Context) error {
	// Initialize git repository
	cmd := exec.Command("git", "init")
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error while initializing git repo: %v", err)
	}
	var configFile []byte
	output := cliCtx.String("output")
	if output != "" {
		configName = output
	}
	if cliCtx.Bool("default") {
		configFile, err = hub.Fs.ReadFile("templates/config.sample.yml")
		if err != nil {
			return fmt.Errorf("error reading default config file template: %v", err)
		}
		err = createDefaultConfig(configFile, configName)
		if err != nil {
			return fmt.Errorf("error creating default config: %v", err)
		}
	} else {
		workloads := []models.Workload{}
		workloadsLen := gatherBasicInfo()
		if output == "" {
			configName = gatherOutputFileInfo()
		}
		// Iterate for all workloads
		for i := 0; i < workloadsLen; i++ {
			wd, err := gatherWorkloadsInfo()
			if err != nil {
				return fmt.Errorf("error while preparing resources for deployment: %v", err)
			}
			workloads = append(workloads, wd)
		}
		var cfg = models.Config{
			Workloads: workloads,
		}
		configFile, err := yaml.Marshal(cfg)
		if err != nil {
			return fmt.Errorf("Error while marshalling yaml: %v", err)
		}
		err = createDefaultConfig(configFile, configName)
		if err != nil {
			return fmt.Errorf("error creating default config: %v", err)
		}
	}
	log.Printf("Congrats! Your default configuration is created at %s", configName)
	return nil
}
