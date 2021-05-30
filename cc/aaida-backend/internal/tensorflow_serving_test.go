package internal_test

import (
	"testing"

	. "github.com/b21-cap0029/bangkit-capstone/cc/aaida-backend/internal"
	"github.com/stretchr/testify/assert"
)

func TestNewDefaultTFServingPredictorPanic(t *testing.T) {
	TensorflowServingHost = ""
	assert.Panics(t, func() { NewDefaultTFServingPredictor() }, "the code didn't panic")
}
