package middleware

import (
	"github.com/gin-gonic/gin"
)

func AdminOnly() gin.HandlerFunc {
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
