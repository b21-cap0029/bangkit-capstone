package internal

// Predictor
type Predictor interface {
	Predict([]string) ([]float32, error)
}
