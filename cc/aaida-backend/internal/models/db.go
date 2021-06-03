package models

import (
	"fmt"
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

func FindCasesWithTwitterUserID(db *gorm.DB, twitterUserID int64) []Case {
	var cases []Case
	db.Where("twitter_user_id = ?", twitterUserID).Find(&cases)
	return cases
}

func FindUserWithLeastUnclosedClaim(db *gorm.DB) (User, error) {
	var user User
	type Result struct {
		ID    string
		Count int
	}
	var results []Result

	tx := db.Model(&User{})
	tx = tx.Select("users.id, count(cases.id) as count")
	tx = tx.Joins(`left join "cases" on users.id = cases.owner_id`)
	tx = tx.Group("users.id").Order("count asc").Scan(&results)
	if tx.Error != nil {
		return User{}, tx.Error
	}

	if len(results) == 0 {
		return User{}, fmt.Errorf("record not found")
	}

	tx = db.Where("id = ?", results[0].ID).First(&user)
	if tx.Error != nil {
		return User{}, tx.Error
	}

	return user, nil
}
