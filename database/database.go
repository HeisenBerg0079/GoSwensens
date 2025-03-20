package database

import (
	"log"

	m "earth/swensen/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Global DB instance
var DB *gorm.DB

// InitializeDB initializes the database connection
func InitializeDB() {
	var err error
	dsn := "root@tcp(127.0.0.1:3306)/go_user?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	// Auto-migrate models (including Products)
	if err := DB.AutoMigrate(&m.Users{}, &m.Products{}); err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	log.Println("Database connected and migrated successfully!")
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return DB
}
