package internal

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var (
	TwitterClientId     = os.Getenv("TWITTER_CLIENT_ID")
	TwitterClientSecret = os.Getenv("TWITTER_CLIENT_SECRET")
	TensorflowBaseURL   = os.Getenv("TENSORFLOW_BASE_URL")
)
