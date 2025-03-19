package main

import (
	"log"
	"main/cmd/server"
	_ "main/docs"
	"main/pkg/db"
	"main/pkg/db/migrations"
	"main/pkg/logger"
	"main/pkg/utils/seeds"

	"gorm.io/gorm"
)

// @title           Dvd Rental API
// @version         1.0
// @description     This is an API for working with the public database **dvd_rental**. \nSupports CRUD operations on movies, actors, rentals, etc. \n**Technologies:**  Go (Gin) + PostgreSQL \n**Database:** dvd_rental \n**Main features:** - Get a list of movies - Search for actors - Make a rental
// @termsOfService  http://example.com/terms/

// @contact.name   Roman S.
// @contact.url    https://www.linkedin.com/in/roman-s-bba6021a5/
// @contact.email  unpredictableanonymous639@gmail.com

// @license.name  GPL-3.0 license
// @license.url   https://www.gnu.org/licenses/gpl-3.0.html

// @host      localhost:8080
// @BasePath  /

func main() {
	// rdb := redisClient.InitRedis()
	db := initialize()
	server.InitServer(db)
}

func initialize() *gorm.DB {
	logger.InfoLogger.Println("Initializing database...")

	database, err := db.InitDb()
	if err != nil {
		logger.ErrorLogger.Printf("Failed to initialize database: %v", err)
	}

	log.Println("Running migrations...")
	// if err := migrations.CreateTables(); err != nil {
	// 	logger.ErrorLogger.Printf("Failed to perform migrations: %v", err)
	// }
	if err := migrations.RunMigrations(); err != nil {
		logger.ErrorLogger.Printf("Failed to perform migrations: %v", err)
	}

	log.Println("Seeding initial data...")
	if err := seeds.SeedLanguageData(); err != nil {
		logger.WarningLogger.Printf("Failed to seed language data: %v", err)
	}

	return database
}
