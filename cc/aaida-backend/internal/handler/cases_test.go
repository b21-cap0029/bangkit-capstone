package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	. "github.com/b21-cap0029/bangkit-capstone/cc/aaida-backend/internal/handler"
	"github.com/b21-cap0029/bangkit-capstone/cc/aaida-backend/internal/models"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestCasesNotFound(t *testing.T) {
	mockDB, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	mockDB.AutoMigrate(&models.User{}, &models.Case{})

	req, err := http.NewRequest("GET", "/cases/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.Handle("/cases/{id:[0-9]+}", NewCasesHandler(mockDB))
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code, "Should be equal")
	assert.Equal(t, "record not found\n", rr.Body.String(), "Should be equal")
}

func TestCasesSuccess(t *testing.T) {
	mockDB, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	mockDB.AutoMigrate(&models.User{}, &models.Case{})
	mockUser := models.User{
		Email:      "giovanmail@gmail.com",
		Name:       "Giovan Isa Musthofa",
		IsVerified: true,
	}
	mockDB.Create(&mockUser)
	mockCase := models.Case{
		TwitterUserID: 1,
		TweetID:       1,
		Class:         "Positive",
		Score:         0.999999,
		Owner:         &mockUser,
		IsClaimed:     true,
		IsClosed:      false,
	}
	mockDB.Create(&mockCase)

	req, err := http.NewRequest("GET", "/cases/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.Handle("/cases/{id:[0-9]+}", NewCasesHandler(mockDB))
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Should be equal")
}
