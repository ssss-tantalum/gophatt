package app

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	HTTPPort  uint64 `env:"HTTP_PORT"`
	SecretKey string `env:"SECRET_KEY"`
	DbDSN     string `env:"DB_DSN"`
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &Config{
		HTTPPort:  cfg.HTTPPort,
		SecretKey: cfg.SecretKey,
		DbDSN:     cfg.DbDSN,
	}, nil
}
