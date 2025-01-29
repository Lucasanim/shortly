package config

import (
	"os"
)

type Config struct {
	Port             string
	BaseUrl          string
	DynamoDbEndpoint string
}

var Env *Config

func LoadEnv() {
	Env = &Config{
		Port:             getEnv("PORT", "3000"),
		BaseUrl:          getEnv("BASE_URL", "http://localhost:3000"),
		DynamoDbEndpoint: getEnv("DYNAMO_DB_ENDPOINT", "http://localhost:8000"),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
