package handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	. "github.com/b21-cap0029/bangkit-capstone/cc/aaida-backend/internal/handler"
	"github.com/b21-cap0029/bangkit-capstone/cc/aaida-backend/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type DBCreatorMock struct {
	mock.Mock
}

func (d *DBCreatorMock) Create(v interface{}) *gorm.DB {
	args := d.Called(v)
	return args.Get(0).(*gorm.DB)
}

func TestCasesSubmitGet(t *testing.T) {
	req, err := http.NewRequest("GET", "/cases/submit", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.Handler(NewCasesSubmitHandler(nil))
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusMethodNotAllowed, rr.Code, "Should be equal")
	assert.Equal(t, "only POST method allowed\n", rr.Body.String(), "Should be equal")
}

func TestCasesSubmitEmptyBody(t *testing.T) {
	req, err := http.NewRequest("POST", "/cases/submit", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.Handler(NewCasesSubmitHandler(nil))
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code, "Should be equal")
	assert.Equal(t, "empty body\n", rr.Body.String(), "Should be equal")
}

func TestCasesSubmit(t *testing.T) {
	caseObj := models.Case{
		TwitterUserID: 1,
		TweetID:       1,
		Class:         "Positive",
		Score:         0.999999,
		IsClaimed:     false,
		IsClosed:      false,
	}

	b, err := json.Marshal(caseObj)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/cases/submit", bytes.NewReader(b))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	mockDB := new(DBCreatorMock)
	now := time.Now().Truncate(time.Nanosecond)
	mockDB.On("Create", &caseObj).Return(&gorm.DB{}).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*models.Case)
		arg.ID = 1
		arg.CreatedDate = now
	})

	rr := httptest.NewRecorder()
	handler := http.Handler(NewCasesSubmitHandler(mockDB))
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Should be equal")

	var returnedCaseObj models.Case
	json.Unmarshal(rr.Body.Bytes(), &returnedCaseObj)

	assert.Equal(t, returnedCaseObj.ID, uint(1), "Should be equal")
	assert.Equal(t, returnedCaseObj.CreatedDate, now, "Should be equal")
}

func TestCasesSubmitDBError(t *testing.T) {
	caseObj := models.Case{
		TwitterUserID: 1,
		TweetID:       1,
		Class:         "Positive",
		Score:         0.999999,
		IsClaimed:     false,
		IsClosed:      false,
	}

	b, err := json.Marshal(caseObj)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/cases/submit", bytes.NewReader(b))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	mockDB := new(DBCreatorMock)
	mockDB.On("Create", &caseObj).Return(&gorm.DB{
		Error: fmt.Errorf("UNIQUE constraint failed: cases.tweet_id"),
	})

	rr := httptest.NewRecorder()
	handler := http.Handler(NewCasesSubmitHandler(mockDB))
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code, "Should be equal")
	assert.Equal(t, "UNIQUE constraint failed: cases.tweet_id\n", rr.Body.String(), "Should be equal")
}
