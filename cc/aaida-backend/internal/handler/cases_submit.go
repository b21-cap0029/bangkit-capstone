package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gorm.io/gorm"

	"github.com/b21-cap0029/bangkit-capstone/cc/aaida-backend/internal/models"
)

type DBCreator interface {
	Create(interface{}) *gorm.DB
}

type CasesSubmitHandler struct {
	db DBCreator
}

func NewDefaultCasesSubmitHandler() *CasesSubmitHandler {
	return NewCasesSubmitHandler(models.DB)
}

func NewCasesSubmitHandler(db DBCreator) *CasesSubmitHandler {
	return &CasesSubmitHandler{db: db}
}

func (c *CasesSubmitHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	var caseObj models.Case

	w.Header().Set("Content-Type", "application/json")
	jsonEnc := json.NewEncoder(w)

	if r.Method != "POST" {
		err = fmt.Errorf("only POST method allowed")
		http.Error(w, err.Error(), http.StatusMethodNotAllowed)
		return
	}

	if r.Body == nil {
		err = fmt.Errorf("empty body")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&caseObj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := c.db.Create(&caseObj)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
		return
	}

	// TODO Matchmaking and notify user

	jsonEnc.Encode(caseObj)
}
