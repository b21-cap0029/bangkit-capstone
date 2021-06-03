package models_test

import (
	"testing"

	. "github.com/b21-cap0029/bangkit-capstone/cc/aaida-backend/internal/models"
	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestFindCasesWithTwitterUserIDDoesntFindAnyUser(t *testing.T) {
	mockDB, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	mockDB.AutoMigrate(&User{}, &Case{})

	cases := FindCasesWithTwitterUserID(mockDB, int64(1))

	assert.Equal(t, []Case{}, cases, "Should be equal")
}

func TestFindCasesWithTwitterUserIDFindTwo(t *testing.T) {
	mockDB, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	err = mockDB.AutoMigrate(&User{}, &Case{})
	if err != nil {
		t.Fatal(err)
	}

	tx := mockDB.Create(&[]Case{
		{TwitterUserID: 1, TweetID: 1, Class: "Positive", Score: 0.6},
		{TwitterUserID: 1, TweetID: 2, Class: "Positive", Score: 0.8},
	})
	if tx.Error != nil {
		t.Fatal(tx.Error)
	}

	cases := FindCasesWithTwitterUserID(mockDB, int64(1))

	assert.Equal(t, 2, len(cases), "Should be equal")
}

func TestFindUserWithLeastUnclosedClaimNoUser(t *testing.T) {
	mockDB, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	err = mockDB.AutoMigrate(&User{}, &Case{})
	if err != nil {
		t.Fatal(err)
	}

	_, err = FindUserWithLeastUnclosedClaim(mockDB)
	assert.EqualError(t, err, "record not found", "Should be equal")
}

func TestFindUserWithLeastUnclosedClaimNoClaimedCase(t *testing.T) {
	mockDB, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	err = mockDB.AutoMigrate(&User{}, &Case{})
	if err != nil {
		t.Fatal(err)
	}

	userMock := User{
		Email:      "giovanmail@gmail.com",
		Name:       "Giovan Isa Musthofa",
		IsVerified: true,
	}
	tx := mockDB.Create(&userMock)
	if tx.Error != nil {
		t.Fatal(tx.Error)
	}

	tx = mockDB.Create(&[]Case{
		{TwitterUserID: 1, TweetID: 1, Class: "Positive", Score: 0.6},
		{TwitterUserID: 1, TweetID: 2, Class: "Positive", Score: 0.8},
	})
	if tx.Error != nil {
		t.Fatal(tx.Error)
	}

	user, err := FindUserWithLeastUnclosedClaim(mockDB)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, userMock, user, "Should be equal")
}

func TestFindUserWithLeastUnclosedClaimFound(t *testing.T) {
	mockDB, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	err = mockDB.AutoMigrate(&User{}, &Case{})
	if err != nil {
		t.Fatal(err)
	}

	userMock := User{
		Email:      "giovanmail@gmail.com",
		Name:       "Giovan Isa Musthofa",
		IsVerified: true,
	}
	tx := mockDB.Create(&userMock)
	if tx.Error != nil {
		t.Fatal(tx.Error)
	}

	tx = mockDB.Create(&[]Case{
		{TwitterUserID: 1, TweetID: 1, Class: "Positive", Score: 0.6, OwnerID: null.IntFrom(userMock.ID)},
		{TwitterUserID: 1, TweetID: 2, Class: "Positive", Score: 0.8},
	})
	if tx.Error != nil {
		t.Fatal(tx.Error)
	}

	user, err := FindUserWithLeastUnclosedClaim(mockDB)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, userMock, user, "Should be equal")
}

func TestFindUserWithLeastUnclosedClaimFoundLeast(t *testing.T) {
	mockDB, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	err = mockDB.AutoMigrate(&User{}, &Case{})
	if err != nil {
		t.Fatal(err)
	}

	usersMock := []User{
		{Email: "giovanmail@gmail.com", Name: "Giovan Isa Musthofa", IsVerified: true},
		{Email: "dzakyale-kampus@yahoo.com", Name: "dzakyraffy", IsVerified: true},
	}
	tx := mockDB.Create(&usersMock)
	if tx.Error != nil {
		t.Fatal(tx.Error)
	}

	tx = mockDB.Create(&[]Case{
		{TwitterUserID: 1, TweetID: 1, Class: "Positive", Score: 0.6, OwnerID: null.IntFrom(usersMock[0].ID)},
		{TwitterUserID: 2, TweetID: 2, Class: "Positive", Score: 0.8, OwnerID: null.IntFrom(usersMock[1].ID)},
		{TwitterUserID: 3, TweetID: 3, Class: "Positive", Score: 0.9, OwnerID: null.IntFrom(usersMock[0].ID)},
	})
	if tx.Error != nil {
		t.Fatal(tx.Error)
	}

	user, err := FindUserWithLeastUnclosedClaim(mockDB)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, usersMock[1], user, "Should be equal")
}
