package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/b21-cap0029/bangkit-capstone/cc/aaida-backend/internal/handler"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type PredictorMock struct {
	mock.Mock
}

func (m *PredictorMock) Predict(texts []string) ([]float32, error) {
	args := m.Called(texts)
	return args.Get(0).([]float32), args.Error(1)
}

type StatusServiceMock struct {
	mock.Mock
}

func (s *StatusServiceMock) Show(id int64, params *twitter.StatusShowParams) (*twitter.Tweet, *http.Response, error) {
	args := s.Called(id, params)
	return args.Get(0).(*twitter.Tweet), args.Get(1).(*http.Response), args.Error(2)
}

func TestCheckHandler(t *testing.T) {
	// https://twitter.com/idwiki/status/1391607030255816704
	text := `"Usia 25 tahun idealnya seperti apa?"

✅ Tidak menghiraukan standar kesuksesan milik orang lain, karena setiap orang punya standarnya masing-masing. Selama kamu bisa berkembang dan bahagia, itu sudah baik adanya. 🙏

Mau di usia berapa pun, semoga kamu bahagia selalu.`
	mockPredictor := new(PredictorMock)
	mockPredictor.On("Predict", []string{text}).Return([]float32{0.0004786849}, nil)

	mockStatusService := new(StatusServiceMock)
	mockStatusService.On(
		"Show",
		int64(1391607030255816704),
		&twitter.StatusShowParams{TweetMode: "extended"},
	).Return(
		&twitter.Tweet{
			FullText: text,
		},
		&http.Response{},
		nil,
	)

	req, err := http.NewRequest("GET", "/check?url=https%3A%2F%2Ftwitter.com%2Fidwiki%2Fstatus%2F1391607030255816704", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.Handler(NewCheckHandler(mockPredictor, mockStatusService, nil))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	assert.Equal(t, http.StatusOK, rr.Code, "Should be equal")
	assert.Equal(t, "{\"class\":\"Negative\",\"score\":0.0004786849}\n",
		rr.Body.String(), "Should be equal")
}
