package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
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

	expected := `{"class":"Negative","score":0.0004786849}` + "\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
