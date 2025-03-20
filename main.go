package main

import (
	d "earth/swensen/database"
	l "earth/swensen/logic"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Database
	d.InitializeDB()

	// Setup Gin Router
	r := gin.Default()

	// Define routes
	r.POST("/register", l.Register)
	r.POST("/login", l.Login)

	// Product Routes
	r.POST("/products", l.CreateProduct)
	r.GET("/products", l.GetProducts)

	// Start Server
	r.Run(":8080")
}
