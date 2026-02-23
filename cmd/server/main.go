package main

import (
	"ecommerce-api/cmd/server/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Ecommerce API running"})
	})

	routes.SetupRoutes(r)

	r.Run(":8080")
}
