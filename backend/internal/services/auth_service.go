package services

import (
	"database/sql"
	"fmt"

	"ecommerce-api/internal/db"
	"ecommerce-api/internal/models"
	"ecommerce-api/internal/utils"
)

func RegisterUser(user models.User) (models.User, error) {
	d := db.DB()
	// check duplicate
	var exists int
	row := d.QueryRow("SELECT COUNT(1) FROM users WHERE email = ?", user.Email)
	row.Scan(&exists)
	if exists > 0 {
		return models.User{}, fmt.Errorf("email already exists")
	}

	hashedPassword, err := utils.HashPassword(user.PasswordHash)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to hash password")
	}
	user.PasswordHash = hashedPassword

	_, err = d.Exec("INSERT INTO users(id,name,email,password_hash,role,blocked) VALUES(?,?,?,?,?,0)", user.ID, user.Name, user.Email, user.PasswordHash, user.Role)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func Authenticate(email, password string) (models.User, error) {
	d := db.DB()
	var u models.User
	row := d.QueryRow("SELECT id,name,email,password_hash,role,blocked FROM users WHERE email = ?", email)
	if err := row.Scan(&u.ID, &u.Name, &u.Email, &u.PasswordHash, &u.Role, &u.Blocked); err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, fmt.Errorf("invalid credentials")
		}
		return models.User{}, err
	}

	if u.Blocked {
		return models.User{}, fmt.Errorf("user is blocked")
	}

	if !utils.CheckPassword(password, u.PasswordHash) {
		return models.User{}, fmt.Errorf("invalid credentials")
	}

	// do not return password hash
	u.PasswordHash = ""
	return u, nil
}

func GetAllUsers() []models.User {
	d := db.DB()
	rows, err := d.Query("SELECT id,name,email,role,blocked FROM users")
	if err != nil {
		return []models.User{}
	}
	defer rows.Close()
	out := []models.User{}
	for rows.Next() {
		var u models.User
		rows.Scan(&u.ID, &u.Name, &u.Email, &u.Role, &u.Blocked)
		out = append(out, u)
	}
	return out
}

func GetUserByID(id string) (models.User, error) {
	d := db.DB()
	var u models.User
	row := d.QueryRow("SELECT id,name,email,role,blocked FROM users WHERE id = ?", id)
	if err := row.Scan(&u.ID, &u.Name, &u.Email, &u.Role, &u.Blocked); err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, fmt.Errorf("user not found")
		}
		return models.User{}, err
	}
	return u, nil
}

func SetUserBlocked(id string, blocked bool) (models.User, error) {
	d := db.DB()
	_, err := d.Exec("UPDATE users SET blocked = ? WHERE id = ?", boolToInt(blocked), id)
	if err != nil {
		return models.User{}, err
	}
	return GetUserByID(id)
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
