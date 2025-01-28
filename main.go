package main

import (
	"main/cmd/server"
	"main/pkg/db"
)

func main() {
	db.InitDb()
	server.InitServer()
}
