package models

import "gorm.io/gorm"

// Products model
type Products struct {
	gorm.Model
	Name        string  `gorm:"not null"`
	Description string  `gorm:"type:text"`
	Price       float64 `gorm:"not null"`
	Stock       int     `gorm:"not null"`
}
