package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email      string `json:"email" gorm:"not null"`
	Name       string `json:"name" gorm:"not null"`
	IsVerified bool   `json:"is_verified"`
}
