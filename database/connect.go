package database

import (
	"fmt"

	"github.com/bytepass/server/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect to the postgres database.
func Connect(host string, user string, password string, dbname string) {
	// parse connection string
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s", host, user, password, dbname)

	// open database connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed connect to the database: %v", err)
	}

	// --> Auto migrations <--
	// user model
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal("Failed to run auto migration for user model: %v", err)
	}

	// cipher model
	err = db.AutoMigrate(&Cipher{})
	if err != nil {
		log.Fatal("Failed to run auto migration for cipher model: %v", err)
	}
	// --> End auto migrations <--

	log.Info("Connected to database!")

	// move a local database variable to a global variable
	DB = db
}
