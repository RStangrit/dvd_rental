package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port                  string
	GinMode               string
	DSN                   string
	MigrationsSourceURL   string
	MigrationsDatabaseURL string
	SecretKey             string
	RedisAddr             string
	RedisPass             string
	RedisDB               string
	RABBITMQ_USER         string
	RABBITMQ_PASSWORD     string
	ELASTICSEARCH_HOST    string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Warning: No .env file found")
	}

	return &Config{
		Port:                  getEnv("SERVER_PORT", "8080"),
		GinMode:               getEnv("GIN_MODE", "debug"),
		DSN:                   getEnv("DATABASE_URL", "host.docker.internal"),
		MigrationsSourceURL:   getEnv("MIGRATIONS_SOURCE_URL", "pkg/db/migrations/migration_files"),
		MigrationsDatabaseURL: getEnv("MIGRATIONS_DATABASE_URL", "postgres://user:password@host.docker.internal:5432/dvd_rental_v2?sslmode=disable&TimeZone=Asia%2FAlmaty"),
		SecretKey:             getEnv("SECRET_KEY", "null"),
		RedisAddr:             getEnv("REDIS_ADDRESS", "host.docker.internal:6379"),
		RedisPass:             getEnv("REDIS_PASS", ""),
		RedisDB:               getEnv("REDIS_DB", "0"),
		RABBITMQ_USER:         getEnv("RABBITMQ_USER", "guest"),
		RABBITMQ_PASSWORD:     getEnv("RABBITMQ_PASSWORD", "guest"),
		ELASTICSEARCH_HOST:    getEnv("ELASTICSEARCH_HOST", "http://elasticsearch:9200"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
