package models

import "gorm.io/gorm"

// Users model
type Users struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Status   string `gorm:"default:'normal'"`
}
