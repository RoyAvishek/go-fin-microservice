package models

type Permission struct {
	ID          int64  `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string `json:"name" gorm:"unique;not null"`
	Description string `json:"description"`
}
