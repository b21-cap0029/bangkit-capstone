package handler

import (
	"fmt"
	"net/http"

	"github.com/b21-cap0029/bangkit-capstone/cc/aaida-backend/internal/models"
	"gorm.io/gorm"
)

type CasesCloseHandler struct {
	db *gorm.DB
}

func NewDefaultCasesCloseHandler() *CasesCloseHandler {
	return NewCasesCloseHandler(models.DB)
}

func NewCasesCloseHandler(db *gorm.DB) *CasesCloseHandler {
	return &CasesCloseHandler{db: db}
}

func (c *CasesCloseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

	if caseObj.IsClosed {
		err = fmt.Errorf("case has already claimed")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	caseObj.IsClosed = true
	c.db.Save(caseObj)
}
