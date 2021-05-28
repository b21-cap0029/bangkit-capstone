package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// TFServingPredictor
type TFServingPredictor struct {
	host,
	modelName,
	modelVersion string
}

func NewDefaultTFServingPredictor() *TFServingPredictor {
	host := os.Getenv("TENSORFLOW_SERVING_HOST")
	if host == "" {
		log.Panicln("TENSORFLOW_SERVING_HOST is unset")
	}
	return NewTFServingPredictor(host, "model", "1")
}

func NewTFServingPredictor(host, modelName, modelVersion string) *TFServingPredictor {
	return &TFServingPredictor{
		host:         host,
		modelName:    modelName,
		modelVersion: modelVersion,
	}
}

func (t *TFServingPredictor) Predict(texts []string) ([]float32, error) {
	if len(texts) == 0 {
		return nil, fmt.Errorf("texts is empty string array")
	}

	reqBody, err := json.Marshal(map[string][][]string{
		"instances": {texts},
	})
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(
		fmt.Sprintf("http://%s:8501/v%s/models/%s:predict", t.host, t.modelVersion, t.modelName),
		"application/json",
		bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var predicResp map[string][][]float32

	err = json.Unmarshal(body, &predicResp)
	if err != nil {
		return nil, err
	}

	var scores []float32

	if predictions, ok := predicResp["predictions"]; ok {
		scores = predictions[0]
	} else {
		return nil, fmt.Errorf(`no "predictions" key`)
	}

	return scores, nil
}
