package models

type User struct {
	DefaultModel
	Email      string `json:"email" gorm:"not null"`
	Name       string `json:"name" gorm:"not null"`
	IsVerified bool   `json:"is_verified"`
}
