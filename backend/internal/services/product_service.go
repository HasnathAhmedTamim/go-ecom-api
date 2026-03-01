package services

import (
	"database/sql"
	"fmt"

	"ecommerce-api/internal/db"
	"ecommerce-api/internal/models"
)

func GetAllProducts() []models.Product {
	d := db.DB()
	rows, err := d.Query("SELECT id, name, price, stock FROM products")
	if err != nil {
		return []models.Product{}
	}
	defer rows.Close()
	out := []models.Product{}
	for rows.Next() {
		var p models.Product
		rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock)
		out = append(out, p)
	}
	return out
}

func GetProductByID(id string) (models.Product, error) {
	d := db.DB()
	var p models.Product
	row := d.QueryRow("SELECT id, name, price, stock FROM products WHERE id = ?", id)
	if err := row.Scan(&p.ID, &p.Name, &p.Price, &p.Stock); err != nil {
		if err == sql.ErrNoRows {
			return models.Product{}, fmt.Errorf("product not found")
		}
		return models.Product{}, err
	}
	return p, nil
}

func CreateProduct(p models.Product) models.Product {
	d := db.DB()
	d.Exec("INSERT INTO products(id,name,price,stock) VALUES(?,?,?,?)", p.ID, p.Name, p.Price, p.Stock)
	return p
}

func UpdateProduct(id string, updated models.Product) (models.Product, error) {
	d := db.DB()
	res, err := d.Exec("UPDATE products SET name=?,price=?,stock=? WHERE id=?", updated.Name, updated.Price, updated.Stock, id)
	if err != nil {
		return models.Product{}, err
	}
	affected, _ := res.RowsAffected()
	if affected == 0 {
		return models.Product{}, fmt.Errorf("product not found")
	}
	return updated, nil
}

func DeleteProduct(id string) error {
	d := db.DB()
	res, err := d.Exec("DELETE FROM products WHERE id=?", id)
	if err != nil {
		return err
	}
	if ra, _ := res.RowsAffected(); ra == 0 {
		return fmt.Errorf("product not found")
	}
	return nil
}

// DeductStock reduces the stock for a product by qty if available.
func DeductStock(id string, qty int) error {
	d := db.DB()
	tx, err := d.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var stock int
	row := tx.QueryRow("SELECT stock FROM products WHERE id=?", id)
	if err := row.Scan(&stock); err != nil {
		return fmt.Errorf("product not found")
	}
	if stock < qty {
		return fmt.Errorf("insufficient stock")
	}
	_, err = tx.Exec("UPDATE products SET stock=? WHERE id=?", stock-qty, id)
	if err != nil {
		return err
	}
	return tx.Commit()
}

// SearchProducts returns paginated products matching query and price range.
func SearchProducts(query string, page, limit int, minPrice, maxPrice float64) ([]models.Product, int) {
	d := db.DB()
	// Simple implementation: only name LIKE and price filters
	where := "WHERE 1=1"
	args := []interface{}{}
	if query != "" {
		where += " AND LOWER(name) LIKE ?"
		args = append(args, "%"+query+"%")
	}
	if minPrice > 0 {
		where += " AND price >= ?"
		args = append(args, minPrice)
	}
	if maxPrice > 0 {
		where += " AND price <= ?"
		args = append(args, maxPrice)
	}

	// Count total
	countQ := "SELECT COUNT(*) FROM products " + where
	var total int
	row := d.QueryRow(countQ, args...)
	row.Scan(&total)

	if limit <= 0 {
		limit = 10
	}
	if page <= 0 {
		page = 1
	}
	offset := (page - 1) * limit

	q := fmt.Sprintf("SELECT id,name,price,stock FROM products %s LIMIT %d OFFSET %d", where, limit, offset)
	rows, err := d.Query(q, args...)
	if err != nil {
		return []models.Product{}, total
	}
	defer rows.Close()
	out := []models.Product{}
	for rows.Next() {
		var p models.Product
		rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock)
		out = append(out, p)
	}
	return out, total
}
