package internal

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(username, password string) error {
	if len(username) == 0 {
		return errors.New("username cannot be empty")
	}

	if len(password) < 8 {
		return errors.New("password must be at least 8 characters")
	}

	if _, exists := userStore[username]; exists {
		return errors.New("username already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := NewUser(username, string(hashedPassword))

	userStore[username] = user

	return nil
}
