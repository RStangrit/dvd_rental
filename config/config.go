package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string
	GinMode   string
	DSN       string
	SecretKey string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found")
	}

	return &Config{
		Port:      getEnv("SERVER_PORT", "8080"),
		GinMode:   getEnv("GIN_MODE", "debug"),
		DSN:       getEnv("DATABASE_URL", "localhost"),
		SecretKey: getEnv("SECRET_KEY", "null"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
