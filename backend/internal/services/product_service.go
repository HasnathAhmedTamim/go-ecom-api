package services

import (
	"ecommerce-api/internal/models"
	"errors"
	"strings"
)

var products = []models.Product{}

func GetAllProducts() []models.Product {
	return products
}

func GetProductByID(id string) (models.Product, error) {
	for _, p := range products {
		if p.ID == id {
			return p, nil
		}
	}
	return models.Product{}, errors.New("product not found")
}

func CreateProduct(p models.Product) models.Product {
	products = append(products, p)
	return p
}

func UpdateProduct(id string, updated models.Product) (models.Product, error) {
	for i, p := range products {
		if p.ID == id {
			products[i] = updated
			return updated, nil
		}
	}
	return models.Product{}, errors.New("product not found")
}

func DeleteProduct(id string) error {
	for i, p := range products {
		if p.ID == id {
			products = append(products[:i], products[i+1:]...)
			return nil
		}
	}
	return errors.New("product not found")
}

// DeductStock reduces the stock for a product by qty if available.
// Returns error if product not found or insufficient stock.
func DeductStock(id string, qty int) error {
	for i, p := range products {
		if p.ID == id {
			if p.Stock < qty {
				return errors.New("insufficient stock")
			}
			products[i].Stock = p.Stock - qty
			return nil
		}
	}
	return errors.New("product not found")
}

// SearchProducts filters products by name substring (case-insensitive)
// and optional price range, then returns paginated results and total count.
func SearchProducts(query string, page, limit int, minPrice, maxPrice float64) ([]models.Product, int) {
	// normalize
	q := strings.TrimSpace(strings.ToLower(query))

	// filter
	filtered := []models.Product{}
	for _, p := range products {
		// name filter
		if q != "" {
			if !strings.Contains(strings.ToLower(p.Name), q) {
				continue
			}
		}

		// price filters
		if minPrice > 0 && p.Price < minPrice {
			continue
		}
		if maxPrice > 0 && p.Price > maxPrice {
			continue
		}

		filtered = append(filtered, p)
	}

	total := len(filtered)

	// pagination bounds
	if limit <= 0 {
		limit = 10
	}
	if page <= 0 {
		page = 1
	}

	start := (page - 1) * limit
	if start >= total {
		return []models.Product{}, total
	}
	end := start + limit
	if end > total {
		end = total
	}

	return filtered[start:end], total
}
