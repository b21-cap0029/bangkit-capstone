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

func TestProfileSuccess(t *testing.T) {
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

	req, err := http.NewRequest("GET", "/profile", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.Handle("/profile", NewProfileHandler(mockDB))
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Should be equal")
	assert.Equal(t, `{"id":1,"email":"giovanmail@gmail.com","name":"Giovan Isa Musthofa","is_verified":true}`+"\n", rr.Body.String(), "Should be equal")
}
