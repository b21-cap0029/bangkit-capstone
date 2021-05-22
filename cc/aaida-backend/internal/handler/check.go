package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"

	internal "github.com/b21-cap0029/bangkit-capstone/cc/aaida-backend/internal"
	"github.com/dghubble/go-twitter/twitter"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

func Check(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	jsonEnc := json.NewEncoder(w)

	tweetUrlStr := r.URL.Query().Get("url")

	var err error
	if tweetUrlStr == "" {
		err = fmt.Errorf("url parameter is required")
		log.Panicln(err.Error())
	}

	tweetUrl, err := url.Parse(tweetUrlStr)
	if err != nil {
		log.Panicln(err.Error())
	}

	switch tweetUrl.Hostname() {
	case "twitter.com", "mobile.twitter.com":
	case "t.co":
		// TODO try to resolve redirect url
		err = fmt.Errorf("%v is not twitter url", tweetUrl.Hostname())
		log.Panicln(err.Error())
	default:
		err = fmt.Errorf("%v is not twitter url", tweetUrl.Hostname())
		log.Panicln(err.Error())
	}

	re := regexp.MustCompile(`/status/(\d+)`)
	matches := re.FindStringSubmatch(tweetUrl.EscapedPath())

	if len(matches) < 2 {
		err = fmt.Errorf("%v is not a status path", tweetUrl.EscapedPath())
		log.Panicln(err.Error())
	}

	tweetId, err := strconv.ParseInt(matches[1], 10, 64)
	if err != nil {
		log.Panicln(err.Error())
	}

	// oauth2 configures a client that uses app credentials to keep a fresh token
	config := &clientcredentials.Config{
		ClientID:     os.Getenv("TWITTER_CLIENT_ID"),
		ClientSecret: os.Getenv("TWITTER_CLIENT_SECRET"),
		TokenURL:     "https://api.twitter.com/oauth2/token",
	}
	// http.Client will automatically authorize Requests
	httpClient := config.Client(oauth2.NoContext)

	// Twitter client
	client := twitter.NewClient(httpClient)

	tweet, _, err := client.Statuses.Show(tweetId, &twitter.StatusShowParams{
		TweetMode: "extended",
	})
	if err != nil {
		log.Panicln(err.Error())
	}

	reqBody, err := json.Marshal(map[string][][]string{
		"instances": {{tweet.FullText}},
	})
	if err != nil {
		log.Panicln(err.Error())
	}

	resp, err := http.Post(
		fmt.Sprintf("http://%s:8501/v1/models/model:predict", os.Getenv("TENSORFLOW_SERVING_HOST")),
		"application/json",
		bytes.NewBuffer(reqBody))
	if err != nil {
		log.Panicln(err.Error())
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln(err.Error())
	}

	var predicResp map[string][][]float32

	err = json.Unmarshal(body, &predicResp)
	if err != nil {
		log.Panicln(err.Error())
	}

	var score float32

	if predictions, ok := predicResp["predictions"]; ok {
		score = predictions[0][0]
	} else {
		log.Panicln(fmt.Errorf(""))
	}

	var message string
	if score < 0.9 {
		message = "Negative"
	} else {
		message = "Positive"
	}

	jsonEnc.Encode(internal.MLResponse{
		Class: message,
		Score: score,
	})
}
