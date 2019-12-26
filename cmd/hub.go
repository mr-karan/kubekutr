package cmd

import (
	"github.com/knadh/stuffbin"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	models "zerodha.tech/kubekutr/models"
)

// Hub represents the structure for all app wide functions and structs
type Hub struct {
	Logger  *logrus.Logger
	Config  models.Config
	Fs      stuffbin.FileSystem
	Version string
}

func NewHub(logger *logrus.Logger, fs stuffbin.FileSystem, buildVersion string) *Hub {
	hub := &Hub{
		Logger:  logger,
		Fs:      fs,
		Version: buildVersion,
	}
	return hub
}

// initApp acts like a middleware to load app managers with Hub before running any command.
func (hub *Hub) initApp(fn cli.ActionFunc) cli.ActionFunc {
	return func(c *cli.Context) error {
		// Initialize config.
		hub.Config = initConfig(c)
		return fn(c)
	}
}
