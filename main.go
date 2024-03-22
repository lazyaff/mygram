package main

import (
	"final-project/config"
	"final-project/database"
)

func main() {
	Port := ":8080"
	database.StartDB()
	config.StartServer().Run(Port)
}