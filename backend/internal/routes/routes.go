package routes

import (
	"ecommerce-api/internal/handlers"
	"ecommerce-api/internal/middleware"
	"ecommerce-api/internal/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")

	// API root for health checks from frontend
	api.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "eCommerce API running 🚀"})
	})

	api.POST("/auth/register", handlers.Register)
	api.POST("/auth/login", handlers.Login)
	api.GET("/auth/me", middleware.AuthMiddleware(), handlers.Me)

	api.GET("/products", handlers.GetProducts)
	api.GET("/products/:id", handlers.GetProduct)

	// Dev-only: trigger seeding of demo data
	api.POST("/_seed", func(c *gin.Context) {
		services.SeedDemoData()
		c.JSON(200, gin.H{"seeded": true})
	})

	admin := api.Group("/admin")
	admin.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		admin.POST("/products", handlers.CreateProduct)
		admin.PUT("/products/:id", handlers.UpdateProduct)
		admin.DELETE("/products/:id", handlers.DeleteProduct)
		admin.GET("/orders", handlers.GetAllOrders)
		admin.GET("/users", handlers.AdminListUsers)
		admin.PUT("/users/:id/block", handlers.AdminBlockUser)
		admin.PUT("/orders/:id/status", handlers.AdminUpdateOrderStatus)
	}

	user := api.Group("/user")
	user.Use(middleware.AuthMiddleware())
	{
		user.POST("/orders", handlers.CreateOrder)
		user.GET("/orders", handlers.GetUserOrders)
		user.GET("/orders/:id", handlers.GetOrderByID)
		user.PUT("/orders/:id/status", handlers.UserUpdateOrderStatus)
	}
}
