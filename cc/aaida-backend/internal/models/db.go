package models

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDataBase() error {
	var db *gorm.DB
	var err error

	postgresDsn := os.Getenv("POSTGRES_DSN")
	if postgresDsn != "" {
		db, err = gorm.Open(postgres.Open(postgresDsn), &gorm.Config{})
	} else {
		db, err = gorm.Open(sqlite.Open("sqlite3.db"), &gorm.Config{})
	}

	if err != nil {
		return err
	}

	DB = db
	DB.AutoMigrate(&User{}, &Case{})

	return nil
}
