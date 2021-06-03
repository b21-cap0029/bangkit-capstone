package handler

import (
	"fmt"
	"net/http"

	"github.com/b21-cap0029/bangkit-capstone/cc/aaida-backend/internal/models"
	"gorm.io/gorm"
)

type CasesSkipHandler struct {
	db *gorm.DB
}

func NewDefaultCasesSkipHandler() *CasesSkipHandler {
	return NewCasesSkipHandler(models.DB)
}

func NewCasesSkipHandler(db *gorm.DB) *CasesSkipHandler {
	return &CasesSkipHandler{db: db}
}

func (c *CasesSkipHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	if r.Method != "POST" {
		err = fmt.Errorf("only POST method allowed")
		http.Error(w, err.Error(), http.StatusMethodNotAllowed)
		return
	}

	caseObj, err := FindCase(c.db, w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := models.FindUserWithLeastUnclosedClaim(c.db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	caseObj.Owner = &user
	tx := c.db.Save(&caseObj)
	if tx.Error != nil {
		http.Error(w, tx.Error.Error(), http.StatusInternalServerError)
		return
	}

	// TODO Notify
}
