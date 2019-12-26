package cmd

import (
	"log"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/file"
	"github.com/urfave/cli"
	"zerodha.tech/kubekutr/models"
)

// initConfig initializes the app's configuration manager.
func initConfig(c *cli.Context) models.Config {
	var cfg = models.Config{}
	var ko = koanf.New(".")

	if len(c.GlobalStringSlice("config")) == 0 {
		log.Fatal("no --config files specified")
	}
	for _, f := range c.GlobalStringSlice("config") {
		log.Printf("reading config: %s", f)
		if err := ko.Load(file.Provider(f), toml.Parser()); err != nil {
			log.Fatalf("error reading config: %v", err)
		}
	}
	// Read the configuration and load it to internal struct.
	failOnReadConfigErr(ko.Unmarshal("", &cfg))
	return cfg
}

func failOnReadConfigErr(err error) {
	if err != nil {
		log.Fatalf("error reading config: %v.", err)
	}
}
