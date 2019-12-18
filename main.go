package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"zerodha.tech/janus/cmd"
)

var (
	// Version of the build. This is injected at build-time.
	buildVersion = "unknown"
	buildDate    = "unknown"
)

func initLogger(verbose bool) *logrus.Logger {
	// Initialize logger
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

func main() {
	// cli.AppHelpTemplate = AppHelpTemplate
	// cli.CommandHelpTemplate = CommandHelpTemplate
	// cli.SubcommandHelpTemplate = SubcommandHelpTemplate
	// Intialize new CLI app
	app := cli.NewApp()
	app.Name = "janus"
	app.Usage = "Dead easy deployment tool"
	app.Version = buildVersion
	app.Author = "Karan Sharma"
	app.Email = "hello@mrkaran.dev"
	// Register command line args.
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "verbose",
			Usage: "Enable verbose logging",
		},
		cli.StringSliceFlag{
			Name: "config, c", Value: &cli.StringSlice{}, Usage: "path to one or more Nest TOML config files", TakesFile: true},
	}
	var (
		logger = initLogger(true)
	)
	// Initialize hub.
	hub := cmd.NewHub(logger, buildVersion)
	// Register commands.
	app.Commands = []cli.Command{
		hub.ScaffoldProject(hub.Config),
		hub.CreateResource(hub.Config),
	}
	// Run the app.
	hub.Logger.Info("Starting janus...")
	err := app.Run(os.Args)
	if err != nil {
		logger.Errorf("Something terrbily went wrong: %s", err)
	}
}
