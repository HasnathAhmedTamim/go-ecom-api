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

// Update Product Handler
func UpdateProduct(c *gin.Context) {
	productID := c.Param("id")
	var updateData models.Product

	if err := c.BindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, p := range products {
		if p.ID == productID {
			// Update fields
			if updateData.Name != "" {
				products[i].Name = updateData.Name
			}
			if updateData.Price != 0 {
				products[i].Price = updateData.Price
			}
			if updateData.Stock != 0 {
				products[i].Stock = updateData.Stock
			}

			c.JSON(http.StatusOK, products[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
}

// Delete Product Handler
func DeleteProduct(c *gin.Context) {
	productID := c.Param("id")

	for i, p := range products {
		if p.ID == productID {
			// Remove product
			products = append(products[:i], products[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
}
