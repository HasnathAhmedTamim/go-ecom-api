package middleware

import (
	"strings"

	"ecommerce-api/internal/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Authorization header missing"})
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := utils.ValidateJWT(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid token"})
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("role", claims.Role)
		c.Next()
	}
}

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		roleVal, exists := c.Get("role")
		roleStr, ok := roleVal.(string)
		if !exists || !ok || roleStr != "admin" {
			c.AbortWithStatusJSON(403, gin.H{"error": "Admin access required"})
			return
		}
		c.Next()
	}
}
