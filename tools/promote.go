package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

func main() {
	dbPath := "data.db"
	d, err := sql.Open("sqlite", dbPath)
	if err != nil {
		log.Fatalf("open db: %v", err)
	}
	defer d.Close()

	res, err := d.Exec("UPDATE users SET role = 'admin' WHERE email = ?", "admin@local.com")
	if err != nil {
		log.Fatalf("update error: %v", err)
	}
	n, _ := res.RowsAffected()
	fmt.Printf("Rows affected: %d\n", n)

	row := d.QueryRow("SELECT id,name,email,role,blocked FROM users WHERE email = ?", "admin@local.com")
	var id, name, email, role string
	var blocked int
	if err := row.Scan(&id, &name, &email, &role, &blocked); err != nil {
		log.Fatalf("select error: %v", err)
	}
	fmt.Printf("User: id=%s name=%s email=%s role=%s blocked=%d\n", id, name, email, role, blocked)
}
