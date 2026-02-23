package handlers

import (
	"net/http"

	"ecommerce-api/cmd/server/internal/models"

	"github.com/gin-gonic/gin"
)

var products = []models.Product{}

// Create Product
func CreateProduct(c *gin.Context) {
	var product models.Product

	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	products = append(products, product)

	c.JSON(http.StatusCreated, product)
}

// Get All Products
func GetProducts(c *gin.Context) {
	c.JSON(http.StatusOK, products)
}

// Get Product by ID
func GetProduct(c *gin.Context) {
	id := c.Param("id")

	for _, product := range products {
		if product.ID == id {
			c.JSON(http.StatusOK, product)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
}
