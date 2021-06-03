package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/b21-cap0029/bangkit-capstone/cc/aaida-backend/internal/models"
	"github.com/gorilla/mux"
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

	vars := mux.Vars(r)
	var caseObj models.Case

	idStr, exist := vars["id"]

	if !exist {
		err = fmt.Errorf("id parameter is required")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tx := c.db.Where("id = ?", id).First(&caseObj)

	if tx.Error != nil {
		http.Error(w, tx.Error.Error(), http.StatusBadRequest)
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
