package db

import (
	"database/sql"
	"log"
	"path/filepath"

	_ "modernc.org/sqlite"
)

var conn *sql.DB

// Init opens (or creates) the sqlite DB and runs migrations.
func Init(path string) (*sql.DB, error) {
	if path == "" {
		path = "data.db"
	}
	p := filepath.Clean(path)
	d, err := sql.Open("sqlite", p)
	if err != nil {
		return nil, err
	}
	conn = d

	if err := migrate(); err != nil {
		return nil, err
	}

	return conn, nil
}

func DB() *sql.DB {
	return conn
}

func migrate() error {
	stmts := []string{
		`CREATE TABLE IF NOT EXISTS users (
            id TEXT PRIMARY KEY,
            name TEXT,
            email TEXT UNIQUE,
            password_hash TEXT,
            role TEXT,
            blocked INTEGER DEFAULT 0
        );`,
		`CREATE TABLE IF NOT EXISTS products (
			id TEXT PRIMARY KEY,
			name TEXT,
			price REAL,
			stock INTEGER,
			image TEXT,
			description TEXT,
			category TEXT,
			brand TEXT
		);`,
		`CREATE TABLE IF NOT EXISTS orders (
            id TEXT PRIMARY KEY,
            user_id TEXT,
            total REAL,
            status TEXT,
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP
        );`,
		`CREATE TABLE IF NOT EXISTS order_items (
            order_id TEXT,
            product_id TEXT,
            qty INTEGER,
            price REAL
        );`,
	}

	for _, s := range stmts {
		if _, err := conn.Exec(s); err != nil {
			log.Println("migration error:", err)
			return err
		}
	}
	return nil
}
