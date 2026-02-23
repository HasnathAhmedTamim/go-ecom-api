package handlers

import (
	"net/http"

	"ecommerce-api/cmd/server/internal/models"

	"github.com/gin-gonic/gin"
)

var orders = []models.Order{}

// Create Order
func CreateOrder(c *gin.Context) {
	var order models.Order

	if err := c.BindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get user_id from JWT
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in token"})
		return
	}
	order.UserID = userID.(string)
	order.Status = "pending"

	// Check product availability and reduce stock
	for _, prodID := range order.Products {
		found := false
		for i, p := range products {
			if p.ID == prodID {
				found = true
				if p.Stock <= 0 {
					c.JSON(http.StatusBadRequest, gin.H{"error": "Product out of stock: " + p.Name})
					return
				}
				// Reduce stock
				products[i].Stock -= 1
			}
		}
		if !found {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Product not found: " + prodID})
			return
		}
	}

	orders = append(orders, order)

	c.JSON(http.StatusCreated, order)
}

// Get All Orders
func GetOrders(c *gin.Context) {
	c.JSON(http.StatusOK, orders)
}
