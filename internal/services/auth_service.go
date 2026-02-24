package services

import (
	"errors"

	"ecommerce-api/internal/models"

	"golang.org/x/crypto/bcrypt"
)

var users = []models.User{}

func RegisterUser(user models.User) (models.User, error) {

	// Check duplicate email
	for _, u := range users {
		if u.Email == user.Email {
			return models.User{}, errors.New("email already exists")
		}
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.PasswordHash),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return models.User{}, errors.New("failed to hash password")
	}

	user.PasswordHash = string(hashedPassword)

	users = append(users, user)
	return user, nil
}

func Authenticate(email, password string) (models.User, error) {

	for _, u := range users {
		if u.Email == email {

			err := bcrypt.CompareHashAndPassword(
				[]byte(u.PasswordHash),
				[]byte(password),
			)

			if err != nil {
				return models.User{}, errors.New("invalid credentials")
			}

			return u, nil
		}
	}

	return models.User{}, errors.New("invalid credentials")
}
