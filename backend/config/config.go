package config

import (
	"url-shortener/pkg/storage/postgres"
)

type Config struct {
	Host     string `config:"HOST" yaml:"host" validate:"required"`
	Port     string `config:"PORT" yaml:"port" validate:"required"`
	Database postgres.Config
}
