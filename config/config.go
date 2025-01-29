package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	Port             string
	BaseUrl          string
	DynamoDbEndpoint string
	RedisAddr        string
	RedisPassword    string
	RedisDb          int
}

var Env *Config

func LoadEnv() {
	Env = &Config{
		Port:             getEnv("PORT", "3000"),
		BaseUrl:          getEnv("BASE_URL", "http://localhost:3000"),
		DynamoDbEndpoint: getEnv("DYNAMO_DB_ENDPOINT", "http://localhost:8000"),
		RedisAddr:        getEnv("REDIS_ADDR", "localhost:6379"),
		RedisPassword:    getEnv("REDIS_PASSWORD", ""),
		RedisDb:          getIntEnv("REDIS_DB", "0"),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func getIntEnv(key, fallback string) int {
	config := getEnv(key, fallback)

	value, err := strconv.Atoi(config)
	if err != nil {
		log.Fatalf("Error parsing env variable: %v", err)
	}

	return value
}
