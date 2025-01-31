package main

import (
	"log"
	"main/cmd/server"
	"main/pkg/db"
	"main/pkg/utils/seeds"
)

func main() {
	initialize()
	server.InitServer()
}

func initialize() {
	log.Println("Initializing database...")
	if err := db.InitDb(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	log.Println("Seeding initial data...")
	if err := seeds.SeedLanguageData(); err != nil {
		log.Fatalf("Failed to seed language data: %v", err)
	}
}
