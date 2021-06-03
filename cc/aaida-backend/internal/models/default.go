package models

type DefaultModel struct {
	ID int64 `json:"id" gorm:"primary_key"`
}
