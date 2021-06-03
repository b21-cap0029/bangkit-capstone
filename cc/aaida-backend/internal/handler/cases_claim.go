package handler

import (
	"fmt"
	"net/http"

	"github.com/b21-cap0029/bangkit-capstone/cc/aaida-backend/internal/models"
	"gorm.io/gorm"
)

type CasesClaimHandler struct {
	db *gorm.DB
}

func NewDefaultCasesClaimHandler() *CasesClaimHandler {
	return NewCasesClaimHandler(models.DB)
}

func NewCasesClaimHandler(db *gorm.DB) *CasesClaimHandler {
	return &CasesClaimHandler{db: db}
}

func (c *CasesClaimHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

	if caseObj.IsClaimed {
		err = fmt.Errorf("case has already claimed")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	caseObj.IsClaimed = true
	c.db.Save(caseObj)
}
