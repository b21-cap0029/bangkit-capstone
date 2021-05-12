package handler

import (
	"encoding/json"
	"net/http"

	internal "github.com/b21-cap0029/bangkit-capstone/cc/internal"
)

func Check(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	tweet := internal.Tweet{
		Id:        1050118621198921728,
		IdStr:     "1050118621198921728",
		CreatedAt: "Wed Oct 10 20:19:24 +0000 2018",
		Text:      "To make room for more expression, we will now count all emojis as equal—including those with gender‍‍‍ ‍‍and skin t… https://t.co/MkGjXf9aXm",
		User: internal.User{
			Id:         6253282,
			IdStr:      "6253282",
			Name:       "Twitter API",
			ScreenName: "twitterapi",
		},
	}

	jsonEnc := json.NewEncoder(w)
	result := CheckUsingMLModel(tweet)
	jsonEnc.Encode(result)
}

func CheckUsingMLModel(tweet internal.Tweet) internal.MLResponse {
	// Magic ML stuff here
	return internal.MLResponse{
		Class: "Normal",
		Score: 0.99,
	}
}
