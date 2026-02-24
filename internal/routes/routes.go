package routes

import (
	"ecommerce-api/internal/handlers"
	"ecommerce-api/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")

	api.POST("/auth/register", handlers.Register)
	api.POST("/auth/login", handlers.Login)

	api.GET("/products", handlers.GetProducts)
	api.GET("/products/:id", handlers.GetProduct)

	admin := api.Group("/admin")
	admin.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		admin.POST("/products", handlers.CreateProduct)
		admin.PUT("/products/:id", handlers.UpdateProduct)
		admin.DELETE("/products/:id", handlers.DeleteProduct)
		admin.GET("/orders", handlers.GetAllOrders)
	}

	user := api.Group("/user")
	user.Use(middleware.AuthMiddleware())
	{
		user.POST("/orders", handlers.CreateOrder)
		user.GET("/orders", handlers.GetUserOrders)
	}
}
