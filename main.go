package main

import (
	"final-project/config"
	"final-project/database"
	"os"
)

func main() {
	Port := envPortOr("8080")
	database.StartDB()
	config.StartServer().Run(Port)
}

func envPortOr(port string) string {
	if envPort := os.Getenv("PORT"); envPort != "" {
	  return ":" + envPort
	}
	return ":" + port
  }