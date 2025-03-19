package main

import (
	"main/cmd/server"
	_ "main/docs"
	"main/pkg/db"
	"main/pkg/db/migrations"
	redisClient "main/pkg/redis"
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
	redisClient := redisClient.InitRedis()
	db, _ := db.InitDb()
	migrations.LaunchMigrations(db)
	server.InitServer(db, redisClient)
}
