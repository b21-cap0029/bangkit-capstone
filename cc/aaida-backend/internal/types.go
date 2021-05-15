package internal

// Tweet
// Ref: https://developer.twitter.com/en/docs/twitter-api/v1/data-dictionary/object-model/tweet
type Tweet struct {
	Id        int64  `json:"id"`
	IdStr     string `json:"id_str"`
	User      User   `json:"user"`
	CreatedAt string `json:"created_at"`
	Text      string `json:"text"`
}

// User
// Ref: https://developer.twitter.com/en/docs/twitter-api/v1/data-dictionary/object-model/user
type User struct {
	Id         int64  `json:"id"`
	IdStr      string `json:"id_str"`
	Name       string `json:"name"`
	ScreenName string `json:"screen_name"`
}

// MLResponse
type MLResponse struct {
	Class string  `json:"class"`
	Score float32 `json:"score"`
}
