package main

import (
	"embed"
	"os"

	"github.com/mr-karan/kubekutr/cmd"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var (
	// Version and date of the build. This is injected at build-time.
	buildVersion = "unknown"
	buildDate    = "unknown"
)

// initLogger initializes logger
func initLogger(verbose bool) *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	// Set logger level
	if verbose {
		logger.SetLevel(logrus.DebugLevel)
		logger.Debug("verbose logging enabled")
	} else {
		logger.SetLevel(logrus.InfoLevel)
	}
	return logger
}

//go:embed templates
var templateFS embed.FS

func main() {
	// Intialize new CLI app
	app := cli.NewApp()
	app.Name = "kubekutr"
	app.Usage = "Cookie cutter for Kubernetes resource manifests"
	app.Version = buildVersion
	app.Author = "Karan Sharma @mrkaran"
	// Register command line args.
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "verbose",
			Usage: "Enable verbose logging",
		},
		cli.StringSliceFlag{
			Name: "config, c", Value: &cli.StringSlice{}, Usage: "path to one or more config files", TakesFile: true},
	}
	var (
		logger = initLogger(true)
	)

	// Initialize hub.
	hub := cmd.NewHub(logger, templateFS, buildVersion)

	// Register commands.
	app.Commands = []cli.Command{
		hub.ScaffoldProject(hub.Config),
		hub.InitProject(hub.Config),
	}
	// Run the app.
	hub.Logger.Info("Starting kubekutr...")
	if err := app.Run(os.Args); err != nil {
		logger.Errorf("Something terrbily went wrong: %s", err)
	}
}
