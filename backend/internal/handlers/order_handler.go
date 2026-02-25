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

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order := models.Order{
		ID:       utils.GenerateID(),
		UserID:   c.GetString("user_id"),
		Products: input.Products,
		Status:   "pending",
	}

	created, err := services.CreateOrder(order)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, created)
}

func GetUserOrders(c *gin.Context) {
	userID := c.GetString("user_id")
	c.JSON(http.StatusOK, services.GetOrdersByUser(userID))
}

func GetAllOrders(c *gin.Context) {
	c.JSON(http.StatusOK, services.GetAllOrders())
}

// User can cancel their own order
func UserUpdateOrderStatus(c *gin.Context) {
	id := c.Param("id")
	userID := c.GetString("user_id")

	var input struct {
		Status string `json:"status"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, err := services.UpdateOrderStatus(id, userID, input.Status, false)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updated)
}

// Admin can update any order status
func AdminUpdateOrderStatus(c *gin.Context) {
	id := c.Param("id")
	// admin middleware ensures role is admin; retrieve user_id for audit if needed
	adminID := c.GetString("user_id")

	var input struct {
		Status string `json:"status"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, err := services.UpdateOrderStatus(id, adminID, input.Status, true)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updated)
}
