package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(host string, user string, password string, dbName string) (*gorm.DB, error) {
	// parse the DSN (data source name) to a gorm.Dialector
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s", host, user, password, dbName)

	// connect to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// if the env variable `MICROPASS_DEBUG` is set to true, enable sql debug logger
	if os.Getenv("MICROPASS_DEBUG") == "true" {
		db = db.Debug()
	}

	// * Run auto migrations

	// user model
	err = db.AutoMigrate(&User{})
	if err != nil {
		return nil, fmt.Errorf("failed to migrate user model: %v", err)
	}

	// cipher model
	err = db.AutoMigrate(&Cipher{})
	if err != nil {
		return nil, fmt.Errorf("failed to migrate cipher model: %v", err)
	}

	log.Println("Connected to the database!")

	// * End auto migrations

	// move a local database variable to a global variable
	DB = db

	return db, nil
}
