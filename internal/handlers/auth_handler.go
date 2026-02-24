package handlers

import (
	"net/http"

	"ecommerce-api/internal/models"
	"ecommerce-api/internal/services"
	"ecommerce-api/internal/utils"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var input struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
		Role     string `json:"role"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Role == "" {
		input.Role = "user"
	}

	user := models.User{
		ID:           utils.GenerateID(),
		Name:         input.Name,
		Email:        input.Email,
		PasswordHash: input.Password, // hashed in service
		Role:         input.Role,
	}

	created, err := services.RegisterUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":    created.ID,
		"name":  created.Name,
		"email": created.Email,
		"role":  created.Role,
	})
}

func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := services.Authenticate(input.Email, input.Password)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	token, err := utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		c.JSON(500, gin.H{"error": "could not generate token"})
		return
	}

	c.JSON(200, gin.H{
		"user":  user,
		"token": token,
	})
}
