package handler

import (
	"encoding/json"
	"net/http"

	"github.com/b21-cap0029/bangkit-capstone/cc/aaida-backend/internal/models"
	"gorm.io/gorm"
)

type ProfileMatchedHandler struct {
	db *gorm.DB
}

func NewDefaultProfileMatchedHandler() *ProfileMatchedHandler {
	return NewProfileMatchedHandler(models.DB)
}

func NewProfileMatchedHandler(db *gorm.DB) *ProfileMatchedHandler {
	return &ProfileMatchedHandler{db: db}
}

func (c *ProfileMatchedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var user models.User
	jsonEnc := json.NewEncoder(w)

	// TODO use auth
	tx := c.db.First(&user)
	if tx.Error != nil {
		http.Error(w, tx.Error.Error(), http.StatusInternalServerError)
		return
	}

	var cases []models.Case
	c.db.Where("owner_id = ? AND is_claimed = true", user.ID).Find(&cases)
	if tx.Error != nil {
		http.Error(w, tx.Error.Error(), http.StatusInternalServerError)
		return
	}

	jsonEnc.Encode(cases)
}
