package main

import (
	"os"

	// auto load functions
	_ "github.com/bytepass/server/config/autoload"
	_ "github.com/joho/godotenv/autoload"

	"github.com/bytepass/server/api"
	"github.com/bytepass/server/database"
)

var (
	postgresHost     = os.Getenv("POSTGRES_HOST")
	postgresUser     = os.Getenv("POSTGRES_USER")
	postgresPassword = os.Getenv("POSTGRES_PASSWORD")
	postgresDb       = os.Getenv("POSTGRES_DB")
)

func main() {
	database.Connect(postgresHost, postgresUser, postgresPassword, postgresDb)

	// run API server
	api.Run()
}
