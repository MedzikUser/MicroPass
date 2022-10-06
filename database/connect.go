package database

import (
	"fmt"

	"github.com/bytepass/server/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect to the postgres database
func Connect(host string, user string, password string, dbname string) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s", host, user, password, dbname)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed connect to the database: ", err)
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal("Failed to run auto migration in the database: ", err)
	}

	log.Info("Connected to database!")

	DB = db
}
