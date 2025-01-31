package main

import (
	"main/cmd/server"
	"main/pkg/db"
	"main/pkg/utils/seeds"
)

func main() {
	db.InitDb()
	seeds.SeedLanguageData()
	server.InitServer()
}
