package main

import (
	"mensina-be/database"
	"mensina-be/server"
)

// @title API Mensina
// @version 1.0
// @description API desenvolvida para projeto academico
// @host localhost:5000
// @BasePath /
func main() {
	database.StartDb()
	server := server.NewServer()
	server.Run()
}
