package services

import (
	"ecommerce-api/internal/models"
	"errors"
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
