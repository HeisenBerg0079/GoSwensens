package logic

import (
	"net/http"

	d "earth/swensen/database"
	m "earth/swensen/models"

	"github.com/gin-gonic/gin"
)

// ProductRequest struct
type ProductRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required"`
	Stock       int     `json:"stock" binding:"required"`
}

// CreateProduct handles adding a new product
func CreateProduct(c *gin.Context) {
	var req ProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	db := d.GetDB()

	newProduct := m.Products{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
	}

	if err := db.Create(&newProduct).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add product"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Product added successfully"})
}

// GetProducts handles retrieving all products
func GetProducts(c *gin.Context) {
	db := d.GetDB()

	var products []m.Products
	if err := db.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}

	c.JSON(http.StatusOK, products)
}
