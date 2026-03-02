package handlers

import (
	"net/http"

	"ecommerce-api/internal/models"
	"ecommerce-api/internal/services"
	"ecommerce-api/internal/utils"

	"github.com/gin-gonic/gin"
)

// Checkout creates an order and returns a mock payment URL (for demo)
func Checkout(c *gin.Context) {
	var input struct {
		Items []struct {
			ID  string `json:"id"`
			Qty int    `json:"qty"`
		} `json:"items"`
		Address string `json:"address"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	productsMap := map[string]int{}
	for _, it := range input.Items {
		productsMap[it.ID] += it.Qty
	}

	order := models.Order{
		ID:       utils.GenerateID(),
		UserID:   c.GetString("user_id"),
		Products: productsMap,
		Status:   "pending",
	}

	created, err := services.CreateOrder(order)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Mock payment URL (would be replaced with real provider in production)
	paymentURL := "https://pay.mock/checkout/" + created.ID

	c.JSON(http.StatusCreated, gin.H{"payment_url": paymentURL, "order": created, "address": input.Address})
}
