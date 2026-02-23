package routes

import (
	"ecommerce-api/cmd/server/internal/handlers"

	"ecommerce-api/cmd/server/internal/middleware"

	"github.com/gin-gonic/gin"
)

// func SetupRoutes(r *gin.Engine) {

// 	api := r.Group("/api")
// 	{
// 		// Public
// 		api.POST("/register", handlers.Register)
// 		api.POST("/login", handlers.Login)

// 		// Products
// 		api.GET("/products", handlers.GetProducts)
// 		api.GET("/products/:id", handlers.GetProduct)

// 		// Protected Routes
// 		protected := api.Group("/")
// 		protected.Use(middleware.AuthMiddleware())
// 		{
// 			protected.POST("/products", handlers.CreateProduct)
// 			protected.POST("/orders", handlers.CreateOrder)
// 			protected.GET("/orders", handlers.GetOrders)
// 		}
// 	}
// }

func SetupRoutes(r *gin.Engine) {

	api := r.Group("/api")
	{
		// Public
		api.POST("/register", handlers.Register)
		api.POST("/login", handlers.Login)
		api.GET("/products", handlers.GetProducts)
		api.GET("/products/:id", handlers.GetProduct)

		// Auth required
		auth := api.Group("/")
		auth.Use(middleware.AuthMiddleware())
		{
			auth.POST("/orders", handlers.CreateOrder)
			auth.GET("/orders", handlers.GetOrders)

			// Admin only group

			admin := auth.Group("/")
			admin.Use(middleware.AdminOnly())
			{
				admin.POST("/products", handlers.CreateProduct)
				admin.PUT("/products/:id", handlers.UpdateProduct)
				admin.DELETE("/products/:id", handlers.DeleteProduct)
			}
		}
	}
}
