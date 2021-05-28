package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/b21-cap0029/bangkit-capstone/cc/aaida-backend/internal/handler"
	"github.com/stretchr/testify/assert"
)

func TestCheckHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/check?url=https%3A%2F%2Ftwitter.com%2Fidwiki%2Fstatus%2F1391607030255816704", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Check)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	assert.Equal(t, http.StatusOK, rr.Code, "Should be equal")
	assert.Equal(t, "{\"class\":\"Negative\",\"score\":0.0004786849}\n",
		rr.Body.String(), "Should be equal")
}
