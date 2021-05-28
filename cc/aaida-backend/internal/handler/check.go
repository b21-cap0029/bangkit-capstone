package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"

	"github.com/b21-cap0029/bangkit-capstone/cc/aaida-backend/internal"
	"github.com/dghubble/go-twitter/twitter"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

type CheckHandler struct {
	predictor     internal.Predictor
	twitterClient *twitter.Client
	httpClient    *http.Client
}

func NewCheckHandler() *CheckHandler {
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

	return &CheckHandler{
		predictor:     internal.NewDefaultTFServingPredictor(),
		twitterClient: client,
		httpClient:    httpClient,
	}
}

func (c *CheckHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

	tweet, _, err := c.twitterClient.Statuses.Show(tweetId, &twitter.StatusShowParams{
		TweetMode: "extended",
	})
	if err != nil {
		log.Panicln(err.Error())
	}

	scores, err := c.predictor.Predict([]string{tweet.FullText})
	if err != nil {
		log.Panicln(err.Error())
	}

	var message string
	if scores[0] < 0.9 {
		message = "Negative"
	} else {
		message = "Positive"
	}

	jsonEnc.Encode(struct {
		Class string  `json:"class"`
		Score float32 `json:"score"`
	}{
		Class: message,
		Score: scores[0],
	})
}
