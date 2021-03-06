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
	caseObj, err := FindCase(c.db, w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jsonEnc := json.NewEncoder(w)
	jsonEnc.Encode(caseObj)
}

func FindCase(db *gorm.DB, w http.ResponseWriter, r *http.Request) (models.Case, error) {
	vars := mux.Vars(r)
	var caseObj models.Case

	idStr, exist := vars["id"]

	if !exist {
		err := fmt.Errorf("id parameter is required")
		return models.Case{}, err
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return models.Case{}, err
	}

	tx := db.Where("id = ?", id).First(&caseObj)

	if tx.Error != nil {
		return models.Case{}, tx.Error
	}

	return caseObj, nil
}
