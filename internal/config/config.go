package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Postgres `yaml:"postgres"`
	HTTP     `yaml:"http"`
}

type Postgres struct {
	Name           string `env-required:"true"  env:"POSTGRES_DB"`
	Password       string `env-required:"true"  env:"POSTGRES_PASSWORD"`
	MaxConnections int    `yaml:"maxConns"`
	Host           string `yaml:"host"`
	Port           string `yaml:"port"`
	Timeout        int    `yaml:"conn_timeout"`
}

type HTTP struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
