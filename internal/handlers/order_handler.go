package handlers

import (
	"net/http"

	"ecommerce-api/internal/models"
	"ecommerce-api/internal/services"
	"ecommerce-api/internal/utils"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	var input struct {
		Products map[string]int `json:"products"`
	}

	c.ShouldBindJSON(&input)

	order := models.Order{
		ID:       utils.GenerateID(),
		UserID:   c.GetString("user_id"),
		Products: input.Products,
		Status:   "pending",
	}

	c.JSON(http.StatusCreated, services.CreateOrder(order))
}

func GetUserOrders(c *gin.Context) {
	userID := c.GetString("user_id")
	c.JSON(http.StatusOK, services.GetOrdersByUser(userID))
}

func GetAllOrders(c *gin.Context) {
	c.JSON(http.StatusOK, services.GetAllOrders())
}
