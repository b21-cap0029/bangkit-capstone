package handler

import (
	"encoding/json"
	"net/http"

	"github.com/b21-cap0029/bangkit-capstone/cc/aaida-backend/internal/models"
	"gorm.io/gorm"
)

type ProfileCasesHandler struct {
	db *gorm.DB
}

func NewDefaultProfileCasesHandler() *ProfileCasesHandler {
	return NewProfileCasesHandler(models.DB)
}

func NewProfileCasesHandler(db *gorm.DB) *ProfileCasesHandler {
	return &ProfileCasesHandler{db: db}
}

func (c *ProfileCasesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var user models.User
	jsonEnc := json.NewEncoder(w)

	// TODO use auth
	tx := c.db.First(&user)
	if tx.Error != nil {
		http.Error(w, tx.Error.Error(), http.StatusInternalServerError)
		return
	}

	var cases []models.Case
	c.db.Where("owner_id = ? AND is_claimed = false", user.ID).Find(&cases)
	if tx.Error != nil {
		http.Error(w, tx.Error.Error(), http.StatusInternalServerError)
		return
	}

	jsonEnc.Encode(cases)
}
