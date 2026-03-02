package models

type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

// Extended fields for storefront
type ProductDetail struct {
	Product
	Image       string `json:"image"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Brand       string `json:"brand"`
}
