package models

import (
	"time"

	"gorm.io/gorm"
)

type Case struct {
	gorm.Model
	CreatedAt     time.Time `json:"created_at"`
	TwitterUserID int64     `json:"twitter_user_id" gorm:"not null"`
	TweetID       int64     `json:"tweet_id" gorm:"unique:not null"`
	Class         string    `json:"class" gorm:"not null"`
	Score         float32   `json:"score" gorm:"not null"`
	OwnerID       uint      `json:"owner_id"`
	Owner         User      `gorm:"foreignKey:OwnerID"`
	IsClaimed     bool      `json:"is_claimed" gorm:"default:false"`
	IsClosed      bool      `json:"is_closed" gorm:"default:false"`
}
