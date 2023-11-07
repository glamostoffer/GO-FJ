package config

import (
	conn "GO-FJ/pkg/postgres_connector"
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	conn.Postgres `yaml:"postgres"`
	HTTP          `yaml:"http"`
}

type HTTP struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	err = cleanenv.ReadConfig("./internal/config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return cfg, nil
}
