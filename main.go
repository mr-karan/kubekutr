package main

import (
	"os"

	"github.com/knadh/stuffbin"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"zerodha.tech/kubekutr/cmd"
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

// initFileSystem initializes the stuffbin FileSystem to provide
// access to bunded static assets to the app.
func initFileSystem(binPath string) (stuffbin.FileSystem, error) {
	fs, err := stuffbin.UnStuff(os.Args[0])
	if err != nil {
		return nil, err
	}
	return fs, nil
}

func main() {
	// cli.AppHelpTemplate = AppHelpTemplate
	// cli.CommandHelpTemplate = CommandHelpTemplate
	// cli.SubcommandHelpTemplate = SubcommandHelpTemplate
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
			Name: "config, c", Value: &cli.StringSlice{}, Usage: "path to one or more Nest TOML config files", TakesFile: true},
	}
	var (
		logger = initLogger(true)
	)
	// Initialize the static file system into which all
	// required static assets (.css, .js files etc.) are loaded.
	fs, err := initFileSystem(os.Args[0])
	if err != nil {
		logger.Errorf("error reading stuffed binary: %v", err)
	}
	// Initialize hub.
	hub := cmd.NewHub(logger, fs, buildVersion)

	// Register commands.
	app.Commands = []cli.Command{
		hub.ScaffoldProject(hub.Config),
	}
	// Run the app.
	hub.Logger.Info("Starting kubekutr...")
	err = app.Run(os.Args)
	if err != nil {
		logger.Errorf("Something terrbily went wrong: %s", err)
	}
}
