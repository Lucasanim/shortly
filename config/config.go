package config

import (
	"os"
)

type Config struct {
	Port    string
	BaseUrl string
}

var Env = &Config{
	Port:    getEnv("PORT", "3000"),
	BaseUrl: getEnv("BASE_URL", "http://localhost:3000"),
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
