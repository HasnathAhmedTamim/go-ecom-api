package services

import (
	"ecommerce-api/internal/models"
	"errors"
)

var orders = []models.Order{}

// CreateOrder checks product stock, deducts quantities, and creates the order.
// Returns error if any product is not found or has insufficient stock.
func CreateOrder(order models.Order) (models.Order, error) {
	// Validate and deduct stock for each product in the order
	for pid, qty := range order.Products {
		if qty <= 0 {
			return models.Order{}, errors.New("invalid quantity for product " + pid)
		}
		if err := DeductStock(pid, qty); err != nil {
			return models.Order{}, err
		}
	}

	orders = append(orders, order)
	return order, nil
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

// UpdateOrderStatus updates the status of an order.
// If isAdmin is false, only the order owner can set status to "cancelled".
func UpdateOrderStatus(id, userID, status string, isAdmin bool) (models.Order, error) {
	// Validate status
	switch status {
	case "pending", "completed", "cancelled":
		// ok
	default:
		return models.Order{}, errors.New("invalid status")
	}

	for i, o := range orders {
		if o.ID == id {
			if !isAdmin {
				if o.UserID != userID {
					return models.Order{}, errors.New("not allowed")
				}
				if status != "cancelled" {
					return models.Order{}, errors.New("users can only cancel orders")
				}
			}

			orders[i].Status = status
			return orders[i], nil
		}
	}

	return models.Order{}, errors.New("order not found")
}
