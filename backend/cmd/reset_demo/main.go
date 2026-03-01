package main

import (
	"database/sql"
	"fmt"
	"log"

	"ecommerce-api/internal/utils"

	_ "modernc.org/sqlite"
)

func main() {
	dbPath := "data.db"
	d, err := sql.Open("sqlite", dbPath)
	if err != nil {
		log.Fatalf("open db: %v", err)
	}
	defer d.Close()

	hashed, err := utils.HashPassword("password")
	if err != nil {
		log.Fatalf("hash password: %v", err)
	}

	res, err := d.Exec("UPDATE users SET password_hash = ? WHERE email LIKE ?", hashed, "demo@local%")
	if err != nil {
		log.Fatalf("update error: %v", err)
	}
	n, _ := res.RowsAffected()
	fmt.Printf("Rows affected: %d\n", n)

	rows, err := d.Query("SELECT id,email,role FROM users WHERE email LIKE ?", "demo@local%")
	if err != nil {
		log.Fatalf("select error: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var id, email, role string
		rows.Scan(&id, &email, &role)
		fmt.Printf("User: id=%s email=%s role=%s\n", id, email, role)
	}

	// ensure email uses .com for consistency with frontend/postman
	if _, err := d.Exec("UPDATE users SET email = ? WHERE email = ?", "demo@local.com", "demo@local"); err == nil {
		fmt.Println("Ensured demo user email is demo@local.com (if existed demo@local)")
	}
}
