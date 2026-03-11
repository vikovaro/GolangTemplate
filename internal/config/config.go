package config

import (
	"os"
)

type Config struct {
	Port        string
	DatabaseURL string
	JWTSecret   string
}

func Load() (*Config, error) {
	cfg := &Config{
		Port:        getEnv("PORT", "8080"),
		DatabaseURL: getEnv("DATABASE_URL", ""),
		JWTSecret:   getEnv("JWT_SECRET", ""),
	}

	return cfg, nil
}

func getEnv(key string, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
