package app

import (
	"file-analizer/config"
	"file-analizer/internal/adapters/controller/cli"
	"file-analizer/internal/usecases"
	"file-analizer/pkg/logger"
	"fmt"
)

type app struct {
	config config.Config
}

// NewApp allows us to load all the configuration to run the application
func NewApp() (*app, error) {
	c, err := config.NewConfig()
	if err != nil {
		return nil, fmt.Errorf("failed loading the configuration %w", err)
	}

	return &app{
		config: c,
	}, nil
}

// Run execute and loads all the depdendencies to execute the process
func (a *app) Run() error {
	loggerFactory := logger.NewFactory(a.config.Log.Level)
	log := loggerFactory.Create()
	log.Info("Application started")

	log.Info("Creating the instance of the diskUsage process")
	diskUsage := usecases.NewDiskUsage(loggerFactory)

	log.Info("Executing the CLI command handler")
	if err := cli.Start(diskUsage, log); err != nil {
		return fmt.Errorf("failed to execute the CLI command handler: %w", err)
	}

	return nil
}
