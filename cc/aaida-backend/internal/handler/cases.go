package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/b21-cap0029/bangkit-capstone/cc/aaida-backend/internal/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type CasesHandler struct {
	db *gorm.DB
}

func NewDefaultCasesHandler() *CasesHandler {
	return NewCasesHandler(models.DB)
}

func NewCasesHandler(db *gorm.DB) *CasesHandler {
	return &CasesHandler{db: db}
}

func (c *CasesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	jsonEnc := json.NewEncoder(w)
	var caseObj models.Case

	idStr, exist := vars["id"]

	if !exist {
		err := fmt.Errorf("id parameter is required")
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

	jsonEnc.Encode(caseObj)
}
