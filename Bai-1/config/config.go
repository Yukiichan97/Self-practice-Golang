package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	PostgreSQL struct {
		Host     string
		Port     int
		User     string
		Password string
		Database string
	}
}

func NewConfig() (*Config, error) {
	var config Config
	if err := envconfig.Process("", &config); err != nil {
		return nil, fmt.Errorf("error processing environment variables: %w", err)
	}

	fmt.Println("port: ", config.PostgreSQL.Port)

	return &config, nil
}
