package services

import (
	"errors"

	"ecommerce-api/internal/models"
	"ecommerce-api/internal/utils"
)

var users = []models.User{}

func RegisterUser(user models.User) (models.User, error) {

	// Check duplicate email
	for _, u := range users {
		if u.Email == user.Email {
			return models.User{}, errors.New("email already exists")
		}
	}

	// Hash password using utils
	hashedPassword, err := utils.HashPassword(user.PasswordHash)
	if err != nil {
		return models.User{}, errors.New("failed to hash password")
	}

	user.PasswordHash = hashedPassword

	users = append(users, user)
	return user, nil
}

func Authenticate(email, password string) (models.User, error) {

	for _, u := range users {
		if u.Email == email {

			if u.Blocked {
				return models.User{}, errors.New("user is blocked")
			}

			if !utils.CheckPassword(password, u.PasswordHash) {
				return models.User{}, errors.New("invalid credentials")
			}

			return u, nil
		}
	}

	return models.User{}, errors.New("invalid credentials")
}

func GetAllUsers() []models.User {
	return users
}

func SetUserBlocked(id string, blocked bool) (models.User, error) {
	for i, u := range users {
		if u.ID == id {
			users[i].Blocked = blocked
			return users[i], nil
		}
	}
	return models.User{}, errors.New("user not found")
}
