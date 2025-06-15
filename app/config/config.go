package config

import (
	"os"
)
type Config struct {
	App *AppConfig
	DB  *DBConfig
}

type DBConfig struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
	Driver	 string
	SSLMode  string
}

type AppConfig struct {
	AppName string
	AppEnv  string
	AppPort string
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	}
	return fallback 
}

func LoadConfig() *Config {
	return &Config{
		App: &AppConfig{
			AppName: getEnv("APP_NAME", ""),
			AppEnv: getEnv("APP_ENV", ""),
			AppPort: getEnv("APP_PORT", ""),
		},

		DB: &DBConfig{
			Host: getEnv("DB_HOST", ""),
			Port: getEnv("DB_PORT", ""),
			Username: getEnv("DB_USER", ""),
			Password: getEnv("DB_PASSWORD", ""),
			Database: getEnv("DB_NAME", ""),
			Driver: getEnv("DB_DRIVER", ""),
			SSLMode: getEnv("DB_SSLMODE", ""),
		},
	}
}
