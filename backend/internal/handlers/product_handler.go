package handlers

import (
	"net/http"
	"strconv"

	"ecommerce-api/internal/models"
	"ecommerce-api/internal/services"
	"ecommerce-api/internal/utils"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	// Query params: q (search), page, limit, min_price, max_price
	q := c.Query("q")
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")
	minPriceStr := c.DefaultQuery("min_price", "0")
	maxPriceStr := c.DefaultQuery("max_price", "0")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}
	minPrice, _ := strconv.ParseFloat(minPriceStr, 64)
	maxPrice, _ := strconv.ParseFloat(maxPriceStr, 64)

	items, total := services.SearchProducts(q, page, limit, minPrice, maxPrice)

	c.JSON(http.StatusOK, gin.H{
		"items": items,
		"page":  page,
		"limit": limit,
		"total": total,
	})
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
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if p.Price <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "price must be greater than 0"})
		return
	}

	if p.Stock < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "stock cannot be negative"})
		return
	}

	p.ID = utils.GenerateID()

	created := services.CreateProduct(p)
	c.JSON(http.StatusCreated, created)
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var p models.Product
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if p.Price <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "price must be greater than 0"})
		return
	}

	if p.Stock < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "stock cannot be negative"})
		return
	}

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
