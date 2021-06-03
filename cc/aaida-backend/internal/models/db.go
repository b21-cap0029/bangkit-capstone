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

	if postgresDsn == "" {
		postgresDsn, _ = TryContructDSNFromCloudRun()
	}

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

func TryContructDSNFromCloudRun() (string, error) {
	var (
		dbUser                 = os.Getenv("DB_USER")                  // e.g. 'my-db-user'
		dbPwd                  = os.Getenv("DB_PASS")                  // e.g. 'my-db-password'
		instanceConnectionName = os.Getenv("INSTANCE_CONNECTION_NAME") // e.g. 'project:region:instance'
		dbName                 = os.Getenv("DB_NAME")                  // e.g. 'my-database'
	)

	if dbUser == "" {
		return "", fmt.Errorf("missing DB_USER")
	} else if dbPwd == "" {
		return "", fmt.Errorf("missing DB_PASS")
	} else if instanceConnectionName == "" {
		return "", fmt.Errorf("missing INSTANCE_CONNECTION_NAME")
	} else if dbName == "" {
		return "", fmt.Errorf("missing DB_NAME")
	}

	socketDir, isSet := os.LookupEnv("DB_SOCKET_DIR")
	if !isSet {
		socketDir = "/cloudsql"
	}

	dbURI := fmt.Sprintf("user=%s password=%s database=%s host=%s/%s", dbUser, dbPwd, dbName, socketDir, instanceConnectionName)
	return dbURI, nil
}
