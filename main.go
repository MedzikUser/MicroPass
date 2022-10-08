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
	postgres_host     = os.Getenv("POSTGRES_HOST")
	postgres_user     = os.Getenv("POSTGRES_USER")
	postgres_password = os.Getenv("POSTGRES_PASSWORD")
	postgres_db       = os.Getenv("POSTGRES_DB")
)

func main() {
	database.Connect(postgres_host, postgres_user, postgres_password, postgres_db)

	// run API server
	api.Run()
}
