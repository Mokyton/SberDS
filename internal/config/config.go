package config

import (
	"SberDS/internal/server"
	"SberDS/internal/store"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	StoreConfig  *store.Config  `envconfig:"DATABASE"`
	ServerConfig *server.Config `envconfig:"SERVER"`
}

func GetConfig() (*Config, error) {
	var cfg = new(Config)

	if err := envconfig.Process("", cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
