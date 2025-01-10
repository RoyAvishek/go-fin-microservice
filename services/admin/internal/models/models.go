package models

import (
	"gorm.io/gorm"
)

func RegisterModels(db *gorm.DB) error {
	return db.AutoMigrate(&User{}, &Role{}, &Permission{})
}
