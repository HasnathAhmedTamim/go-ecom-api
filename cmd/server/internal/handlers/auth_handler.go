package handlers

import (
	"net/http"

	"ecommerce-api/cmd/server/internal/models"
	"ecommerce-api/cmd/server/internal/utils"

	"github.com/gin-gonic/gin"
)

var users = []models.User{}

// Register
func Register(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user.Password = hashedPassword

	if user.Role == "" {
		user.Role = "user"
	}

	users = append(users, user)

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered",
	})
}

// Login
func Login(c *gin.Context) {
	var loginData models.User

	if err := c.BindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, user := range users {
		if user.Email == loginData.Email &&
			utils.CheckPassword(loginData.Password, user.Password) {

			token, err := utils.GenerateToken(user.ID, user.Role)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Token generation failed"})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"message": "Login successful",
				"token":   token,
			})
			return
		}
	}

	c.JSON(http.StatusUnauthorized, gin.H{
		"error": "Invalid credentials",
	})
}
