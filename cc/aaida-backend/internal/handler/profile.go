package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/b21-cap0029/bangkit-capstone/cc/aaida-backend/internal/models"
	"gorm.io/gorm"
)

type ProfileHandler struct {
	db *gorm.DB
}

func NewDefaultProfileHandler() *ProfileHandler {
	return NewProfileHandler(models.DB)
}

func NewProfileHandler(db *gorm.DB) *ProfileHandler {
	return &ProfileHandler{db: db}
}

func (c *ProfileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var err error
	jsonEnc := json.NewEncoder(w)

	// TODO use auth
	tx := c.db.First(&user)
	if tx.Error != nil {
		http.Error(w, tx.Error.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method == "GET" {
		jsonEnc.Encode(user)
	} else if r.Method == "POST" {

		if r.Body == nil {
			err = fmt.Errorf("empty body")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var updatedUser models.User

		err = json.NewDecoder(r.Body).Decode(&updatedUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		c.db.Save(&updatedUser)
		if tx.Error != nil {
			http.Error(w, tx.Error.Error(), http.StatusInternalServerError)
			return
		}
	}
}
