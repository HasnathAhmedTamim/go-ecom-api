package main

import (
	"log"
	"os"

	"ecommerce-api/internal/db"
	"ecommerce-api/internal/routes"
	"ecommerce-api/internal/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	// Ensure JWT secret exists
	if os.Getenv("JWT_SECRET") == "" {
		log.Fatal("JWT_SECRET not set in environment")
	}

	r := gin.Default()
	// Allow CORS so the Vite frontend can call the API during development
	r.Use(cors.Default())
	r.SetTrustedProxies(nil)

	// Root route
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "eCommerce API running 🚀",
		})
	})

	// Ignore favicon
	r.GET("/favicon.ico", func(c *gin.Context) {
		c.Status(204)
	})

	routes.SetupRoutes(r)

	// Initialize DB (creates data.db in backend/)
	if _, err := db.Init(""); err != nil {
		log.Fatal("failed to init db:", err)
	}

	// Seed demo data for local development (will use DB-backed services)
	services.SeedDemoData()

	// Dynamic port for production
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
