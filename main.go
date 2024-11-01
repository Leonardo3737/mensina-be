package main

import (
	"mensina-be/database"
	"mensina-be/server"
)

func main() {
	database.StartDb()
	server := server.NewServer()

	server.Run()
}
