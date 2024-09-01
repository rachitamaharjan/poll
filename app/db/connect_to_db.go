package db

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := os.Getenv("DB_CONFIG")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error connecting to DB")
	}
}

func CloseDBConnection(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		log.Panicf("Failed to get SQL database connection. Error: %v", err)
		return err
	}

	if err := sqlDB.Close(); err != nil {
		log.Panicf("Failed to close database connection. Error: %v", err)
		return err
	}

	return nil
}
