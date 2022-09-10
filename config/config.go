package config

import (
	"fmt"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type (
	Config struct {
		App
		Log
		DataStore
	}

	App struct {
		SystemAcronym string `env:"SYSTEM_ACRONYM, required"`
		Environment   string `env:"ENVIRONMENT, required"`
		Version       string `env:"VERSION, required"`
	}

	Log struct {
		Level string `env:"LOG_LEVEL, required"`
	}

	DataStore struct {
		FileType string `env:"FILE_TYPE, required"`
	}
)

// NewConfig loads the configuration of the application
func NewConfig() (Config, error) {
	godotenv.Load(".env")

	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return Config{}, fmt.Errorf("loading the configuration %w", err)
	}

	return cfg, nil
}
