package models

import (
	"time"
)

type Case struct {
	DefaultModel
	CreatedDate   time.Time `json:"created_date" gorm:"autoCreateTime:not null"`
	TwitterUserID int64     `json:"twitter_user_id" gorm:"not null"`
	TweetID       int64     `json:"tweet_id" gorm:"unique:not null"`
	Class         string    `json:"class" gorm:"not null"`
	Score         float32   `json:"score" gorm:"not null"`
	OwnerID       uint      `json:"owner_id"`
	Owner         *User     `json:"-" gorm:"foreignKey:OwnerID"`
	IsClaimed     bool      `json:"is_claimed" gorm:"default:false"`
	IsClosed      bool      `json:"is_closed" gorm:"default:false"`
}
