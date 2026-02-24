package services

import (
	"ecommerce-api/internal/models"
)

var orders = []models.Order{}

func CreateOrder(order models.Order) models.Order {
	orders = append(orders, order)
	return order
}

func GetOrdersByUser(userID string) []models.Order {
	userOrders := []models.Order{}
	for _, o := range orders {
		if o.UserID == userID {
			userOrders = append(userOrders, o)
		}
	}
	return userOrders
}

func GetAllOrders() []models.Order {
	return orders
}
