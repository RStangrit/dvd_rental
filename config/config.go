package config

import (
	"log"
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
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found")
	}

	return &Config{
		Port:                  getEnv("SERVER_PORT", "8080"),
		GinMode:               getEnv("GIN_MODE", "debug"),
		DSN:                   getEnv("DATABASE_URL", "localhost"),
		MigrationsSourceURL:   getEnv("MIGRATIONS_SOURCE_URL", "pkg/db/migrations/migration_files"),
		MigrationsDatabaseURL: getEnv("MIGRATIONS_DATABASE_URL", "postgres://user:password@localhost:5432/dvd_rental_v2?sslmode=disable&TimeZone=Asia%2FAlmaty"),
		SecretKey:             getEnv("SECRET_KEY", "null"),
		RedisAddr:             getEnv("REDIS_ADDRESS", "localhost:6379"),
		RedisPass:             getEnv("REDIS_PASS", ""),
		RedisDB:               getEnv("REDIS_DB", "0"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
