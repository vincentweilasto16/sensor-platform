package config

import (
	"os"
)

type AppConfig struct {
	AppName string
	AppPort string
}

func LoadAppConfig() *AppConfig {
	return &AppConfig{
		AppName: os.Getenv("APP_NAME"),
		AppPort: os.Getenv("APP_PORT"),
	}
}
