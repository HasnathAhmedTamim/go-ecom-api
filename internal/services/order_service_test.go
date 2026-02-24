package services

import (
	"testing"

	"ecommerce-api/internal/models"
)

func TestUpdateOrderStatus_AdminCanUpdate(t *testing.T) {
	// Reset orders
	orders = []models.Order{}

	o := models.Order{ID: "o1", UserID: "u1", Products: map[string]int{}, Status: "pending"}
	orders = append(orders, o)

	updated, err := UpdateOrderStatus("o1", "admin", "completed", true)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if updated.Status != "completed" {
		t.Fatalf("expected status completed, got %s", updated.Status)
	}
}

func TestUpdateOrderStatus_UserCanCancelOwnOrder(t *testing.T) {
	orders = []models.Order{}
	o := models.Order{ID: "o2", UserID: "u2", Products: map[string]int{}, Status: "pending"}
	orders = append(orders, o)

	updated, err := UpdateOrderStatus("o2", "u2", "cancelled", false)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if updated.Status != "cancelled" {
		t.Fatalf("expected status cancelled, got %s", updated.Status)
	}
}

func TestUpdateOrderStatus_UserCannotCompleteOrder(t *testing.T) {
	orders = []models.Order{}
	o := models.Order{ID: "o3", UserID: "u3", Products: map[string]int{}, Status: "pending"}
	orders = append(orders, o)

	_, err := UpdateOrderStatus("o3", "u3", "completed", false)
	if err == nil {
		t.Fatalf("expected error when user tries to complete order")
	}
}

func TestUpdateOrderStatus_InvalidStatus(t *testing.T) {
	orders = []models.Order{}
	o := models.Order{ID: "o4", UserID: "u4", Products: map[string]int{}, Status: "pending"}
	orders = append(orders, o)

	_, err := UpdateOrderStatus("o4", "u4", "shipped", true)
	if err == nil {
		t.Fatalf("expected error for invalid status")
	}
}
