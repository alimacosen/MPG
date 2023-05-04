package config

import (
	"os"
)

type Config struct {
	DBConnectionString string
}

func New() *Config {
	return &Config{
		DBConnectionString: os.Getenv("DB_CONNECTION_STRING"),
	}
}
