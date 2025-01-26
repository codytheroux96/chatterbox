package internal

import (
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var sessionStore = make(map[string]string)

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

func LoginUser(username, password string) (string, error) {
	user, exists := userStore[username]
	if !exists {
		return "", errors.New("Invalid username or password")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("Invalid username or password")
	}

	token := uuid.NewString()

	sessionStore[token] = username

	return token, nil
}
