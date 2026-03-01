package services

import (
	"database/sql"
	"ecommerce-api/internal/db"
	"ecommerce-api/internal/models"
	"errors"
	"fmt"
)

// CreateOrder stores the order in the DB and deducts stock within a transaction.
func CreateOrder(order models.Order) (models.Order, error) {
	d := db.DB()
	tx, err := d.Begin()
	if err != nil {
		return models.Order{}, err
	}
	defer tx.Rollback()

	total := 0.0
	// validate and deduct stock
	for pid, qty := range order.Products {
		if qty <= 0 {
			return models.Order{}, errors.New("invalid quantity for product " + pid)
		}

		var stock int
		var price float64
		row := tx.QueryRow("SELECT stock, price FROM products WHERE id = ?", pid)
		if err := row.Scan(&stock, &price); err != nil {
			if err == sql.ErrNoRows {
				return models.Order{}, fmt.Errorf("product %s not found", pid)
			}
			return models.Order{}, err
		}
		if stock < qty {
			return models.Order{}, fmt.Errorf("insufficient stock for product %s", pid)
		}

		// deduct stock
		if _, err := tx.Exec("UPDATE products SET stock = stock - ? WHERE id = ?", qty, pid); err != nil {
			return models.Order{}, err
		}

		total += price * float64(qty)
	}

	// insert order
	if _, err := tx.Exec("INSERT INTO orders(id, user_id, total, status) VALUES(?,?,?,?)", order.ID, order.UserID, total, order.Status); err != nil {
		return models.Order{}, err
	}

	// insert items
	for pid, qty := range order.Products {
		var price float64
		if err := tx.QueryRow("SELECT price FROM products WHERE id = ?", pid).Scan(&price); err != nil {
			return models.Order{}, err
		}
		if _, err := tx.Exec("INSERT INTO order_items(order_id, product_id, qty, price) VALUES(?,?,?,?)", order.ID, pid, qty, price); err != nil {
			return models.Order{}, err
		}
	}

	if err := tx.Commit(); err != nil {
		return models.Order{}, err
	}

	order.Total = total
	return order, nil
}

// GetOrdersByUser returns persisted orders for a user.
func GetOrdersByUser(userID string) []models.Order {
	d := db.DB()
	rows, err := d.Query("SELECT id, status, total, created_at FROM orders WHERE user_id = ? ORDER BY created_at DESC", userID)
	if err != nil {
		return []models.Order{}
	}
	defer rows.Close()

	out := []models.Order{}
	for rows.Next() {
		var o models.Order
		var created sql.NullString
		if err := rows.Scan(&o.ID, &o.Status, &o.Total, &created); err != nil {
			continue
		}
		if created.Valid {
			o.CreatedAt = created.String
		}

		// load items
		items, err := d.Query("SELECT product_id, qty FROM order_items WHERE order_id = ?", o.ID)
		if err == nil {
			prod := map[string]int{}
			for items.Next() {
				var pid string
				var qty int
				items.Scan(&pid, &qty)
				prod[pid] = qty
			}
			items.Close()
			o.Products = prod
		}

		out = append(out, o)
	}
	return out
}

func GetOrderByID(id string) (models.Order, error) {
	d := db.DB()
	var o models.Order
	var created sql.NullString
	row := d.QueryRow("SELECT id, user_id, status, total, created_at FROM orders WHERE id = ?", id)
	if err := row.Scan(&o.ID, &o.UserID, &o.Status, &o.Total, &created); err != nil {
		if err == sql.ErrNoRows {
			return models.Order{}, errors.New("order not found")
		}
		return models.Order{}, err
	}
	if created.Valid {
		o.CreatedAt = created.String
	}

	// load items
	items, err := d.Query("SELECT product_id, qty FROM order_items WHERE order_id = ?", o.ID)
	if err == nil {
		prod := map[string]int{}
		for items.Next() {
			var pid string
			var qty int
			items.Scan(&pid, &qty)
			prod[pid] = qty
		}
		items.Close()
		o.Products = prod
	}

	return o, nil
}

func GetAllOrders() []models.Order {
	d := db.DB()
	rows, err := d.Query("SELECT id FROM orders ORDER BY created_at DESC")
	if err != nil {
		return []models.Order{}
	}
	defer rows.Close()
	out := []models.Order{}
	for rows.Next() {
		var id string
		rows.Scan(&id)
		o, err := GetOrderByID(id)
		if err == nil {
			out = append(out, o)
		}
	}
	return out
}

// UpdateOrderStatus updates status and restores stock on cancellation when appropriate.
func UpdateOrderStatus(id, userID, status string, isAdmin bool) (models.Order, error) {
	// Validate status
	switch status {
	case "pending", "completed", "cancelled":
		// ok
	default:
		return models.Order{}, errors.New("invalid status")
	}

	d := db.DB()

	// fetch order
	o, err := GetOrderByID(id)
	if err != nil {
		return models.Order{}, err
	}

	if !isAdmin {
		if o.UserID != userID {
			return models.Order{}, errors.New("not allowed")
		}
		if status != "cancelled" {
			return models.Order{}, errors.New("users can only cancel orders")
		}
	}

	tx, err := d.Begin()
	if err != nil {
		return models.Order{}, err
	}
	defer tx.Rollback()

	// if cancelling, restore stock
	if status == "cancelled" {
		for pid, qty := range o.Products {
			if _, err := tx.Exec("UPDATE products SET stock = stock + ? WHERE id = ?", qty, pid); err != nil {
				return models.Order{}, err
			}
		}
	}

	if _, err := tx.Exec("UPDATE orders SET status = ? WHERE id = ?", status, id); err != nil {
		return models.Order{}, err
	}

	if err := tx.Commit(); err != nil {
		return models.Order{}, err
	}

	// return updated order
	return GetOrderByID(id)
}
