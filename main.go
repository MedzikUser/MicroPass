package main

import (
	"os"

	"github.com/bytepass/server/api"
	"github.com/bytepass/server/database"

	_ "github.com/joho/godotenv/autoload"
)

var (
	postgres_host     = os.Getenv("POSTGRES_HOST")
	postgres_user     = os.Getenv("POSTGRES_USER")
	postgres_password = os.Getenv("POSTGRES_PASSWORD")
	postgres_db       = os.Getenv("POSTGRES_DB")
)

// // Parse configuration file before init modules,
// // configuration file is required by the `database` module.
// func init() {
// 	config.ParseConfig("config.toml")
// }

func main() {
	database.Connect(postgres_host, postgres_user, postgres_password, postgres_db)

	//database.NewUser("medzik@duck.com", "Super#secret#passphrase")

	api.Start()
}
