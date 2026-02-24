package handlers

import (
	"net/http"

	"ecommerce-api/internal/models"
	"ecommerce-api/internal/services"
	"ecommerce-api/internal/utils"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	c.JSON(http.StatusOK, services.GetAllProducts())
}

func GetProduct(c *gin.Context) {
	id := c.Param("id")

	product, err := services.GetProductByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

func CreateProduct(c *gin.Context) {
	var p models.Product
	c.ShouldBindJSON(&p)

	p.ID = utils.GenerateID()

	c.JSON(http.StatusCreated, services.CreateProduct(p))
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var p models.Product
	c.ShouldBindJSON(&p)
	p.ID = id

	updated, err := services.UpdateProduct(id, p)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updated)
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	err := services.DeleteProduct(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product deleted"})
}
