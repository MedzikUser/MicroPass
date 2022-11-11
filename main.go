package main

import (
	"log"
	"os"

	// run autoload functions
	_ "github.com/MedzikUser/MicroPass/utils"
	_ "github.com/joho/godotenv/autoload"

	"github.com/MedzikUser/MicroPass/database"
	"github.com/MedzikUser/MicroPass/server"
)

var (
	postgresHost     = os.Getenv("POSTGRES_HOST")
	postgresUser     = os.Getenv("POSTGRES_USER")
	postgresPassword = os.Getenv("POSTGRES_PASSWORD")
	postgresDb       = os.Getenv("POSTGRES_DB")
)

func main() {
	_, err := database.Connect(postgresHost, postgresUser, postgresPassword, postgresDb)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	server.Run()
}
