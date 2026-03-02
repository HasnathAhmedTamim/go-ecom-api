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
		p1 := models.ProductDetail{Product: models.Product{ID: utils.GenerateID(), Name: "Aurora Wireless Gaming Headset", Price: 129.99, Stock: 40}, Image: "/screenshots/headset.svg", Description: "Immersive 7.1 surround sound with lightweight comfort.", Category: "Audio", Brand: "Aurora"}
		p2 := models.ProductDetail{Product: models.Product{ID: utils.GenerateID(), Name: "Titan Mechanical Keyboard (RGB)", Price: 159.99, Stock: 25}, Image: "/screenshots/keyboard.svg", Description: "Hot-swappable switches and full RGB lighting.", Category: "Keyboards", Brand: "Titan"}
		p3 := models.ProductDetail{Product: models.Product{ID: utils.GenerateID(), Name: "Phantom Precision Gaming Mouse", Price: 79.99, Stock: 60}, Image: "/screenshots/mouse.svg", Description: "High-DPI optical sensor with programmable buttons.", Category: "Mice", Brand: "Phantom"}
		p4 := models.ProductDetail{Product: models.Product{ID: utils.GenerateID(), Name: "Nebula Ergonomic Gaming Chair", Price: 299.99, Stock: 12}, Image: "/screenshots/chair.svg", Description: "Ergonomic support with breathable materials and adjustable lumbar.", Category: "Chairs", Brand: "Nebula"}
		p5 := models.ProductDetail{Product: models.Product{ID: utils.GenerateID(), Name: "Velocity Pro Controller", Price: 69.99, Stock: 50}, Image: "/screenshots/controller.svg", Description: "Low-latency wireless controller with programmable paddles.", Category: "Controllers", Brand: "Velocity"}
		CreateProduct(p1)
		CreateProduct(p2)
		CreateProduct(p3)
		CreateProduct(p4)
		CreateProduct(p5)
		log.Println("seed: added sample products")
	}

	// Ensure existing rows reference SVG placeholders if they used JPG paths previously
	if _, err := d.Exec("UPDATE products SET image = REPLACE(image, '.jpg', '.svg') WHERE image LIKE '%.jpg'"); err == nil {
		log.Println("seed: updated product image paths to .svg where applicable")
	}

	var ucount int
	row = d.QueryRow("SELECT COUNT(1) FROM users")
	row.Scan(&ucount)
	if ucount == 0 {
		demo := models.User{
			ID:           utils.GenerateID(),
			Name:         "Demo User",
			Email:        "demo@local.com",
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
			Email:        "admin@local.com",
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
