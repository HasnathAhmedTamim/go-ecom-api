package main

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "../backend/data.db")
	if err != nil {
		fmt.Println("open error:", err)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, user_id, total, status, created_at FROM orders")
	if err != nil {
		fmt.Println("query error:", err)
		return
	}
	defer rows.Close()
	count := 0
	for rows.Next() {
		var id, uid, status, created sql.NullString
		var total sql.NullFloat64
		if err := rows.Scan(&id, &uid, &total, &status, &created); err != nil {
			fmt.Println("scan error:", err)
			continue
		}
		fmt.Printf("id=%v user=%v total=%v status=%v created=%v\n", id.String, uid.String, total.Float64, status.String, created.String)
		count++
	}
	fmt.Printf("rows: %d\n", count)
}
