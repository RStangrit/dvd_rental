package main

import (
	"log"
	"main/cmd/server"
	"main/pkg/db"
	"main/pkg/db/migrations"
	"main/pkg/logger"
	"main/pkg/utils/seeds"

	"gorm.io/gorm"
)

func main() {
	db := initialize()
	server.InitServer(db)
}

func initialize() *gorm.DB {
	// log.Println("Initializing database...")
	logger.InfoLogger.Println("Initializing database...")

	database, err := db.InitDb()
	if err != nil {
		// log.Fatalf("Failed to initialize database: %v", err)
		logger.ErrorLogger.Printf("Failed to initialize database: %v", err)
	}

	log.Println("Running migrations...")
	if err := migrations.CreateTables(); err != nil {
		// log.Fatalf("Failed to perform migrations: %v", err)
		logger.ErrorLogger.Printf("Failed to perform migrations: %v", err)
	}

	log.Println("Seeding initial data...")
	if err := seeds.SeedLanguageData(); err != nil {
		// log.Fatalf("Failed to seed language data: %v", err)
		logger.WarningLogger.Printf("Failed to seed language data: %v", err)
	}

	return database
}
