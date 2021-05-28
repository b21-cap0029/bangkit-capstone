package internal_test

import (
	"os"
	"testing"

	. "github.com/b21-cap0029/bangkit-capstone/cc/aaida-backend/internal"
)

func TestNewDefaultTFServingPredictorPanic(t *testing.T) {
	os.Setenv("TENSORFLOW_SERVING_HOST", "")
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	NewDefaultTFServingPredictor()
}
