package migrations

import (
	"log"
	"main/pkg/logger"
	"main/pkg/utils/seeds"

	"gorm.io/gorm"
)

func LaunchMigrations(db *gorm.DB) {
	logger.InfoLogger.Println("Initializing database...")

	log.Println("running migrations...")
	// if err := launchMigrationsGORM(); err != nil {
	// 	logger.ErrorLogger.Printf("Failed to perform migrations: %v", err)
	// }
	if err := launchMigrationsMigrate(); err != nil {
		logger.ErrorLogger.Printf("failed to perform migrations: %v", err)
	}

	log.Println("seeding initial data...")
	if err := seeds.SeedLanguageData(); err != nil {
		logger.WarningLogger.Printf("failed to seed language data: %v", err)
	}
}
