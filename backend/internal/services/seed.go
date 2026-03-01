package services

import (
	"log"

	"ecommerce-api/internal/db"
	"ecommerce-api/internal/models"
	"ecommerce-api/internal/utils"
)

// SeedDemoData adds example products and a demo user when the in-memory
// stores are empty. Safe to call multiple times.
func SeedDemoData() {
	d := db.DB()
	// check products count
	var pcount int
	row := d.QueryRow("SELECT COUNT(1) FROM products")
	row.Scan(&pcount)
	if pcount == 0 {
		p1 := models.Product{ID: utils.GenerateID(), Name: "Red T-Shirt", Price: 19.99, Stock: 50}
		p2 := models.Product{ID: utils.GenerateID(), Name: "Blue Jeans", Price: 49.99, Stock: 30}
		p3 := models.Product{ID: utils.GenerateID(), Name: "Sneakers", Price: 79.99, Stock: 20}
		CreateProduct(p1)
		CreateProduct(p2)
		CreateProduct(p3)
		log.Println("seed: added sample products")
	}

	var ucount int
	row = d.QueryRow("SELECT COUNT(1) FROM users")
	row.Scan(&ucount)
	if ucount == 0 {
		demo := models.User{
			ID:           utils.GenerateID(),
			Name:         "Demo User",
			Email:        "demo@local",
			PasswordHash: "password",
			Role:         "user",
		}
		if _, err := RegisterUser(demo); err != nil {
			log.Println("seed: failed to register demo user:", err)
		} else {
			log.Println("seed: created demo user demo@local (password: password)")
		}

		admin := models.User{
			ID:           utils.GenerateID(),
			Name:         "Admin",
			Email:        "admin@local",
			PasswordHash: "adminpass",
			Role:         "admin",
		}
		if _, err := RegisterUser(admin); err != nil {
			log.Println("seed: failed to register admin user:", err)
		} else {
			log.Println("seed: created admin user admin@local (password: adminpass)")
		}
	}
}

