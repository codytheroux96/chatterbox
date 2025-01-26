package internal

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdat"`
}

var userStore = make(map[string]*User)

func NewUser(username, hashedPassword string) *User {
	return &User{
		ID:        uuid.New(),
		Username:  username,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
	}
}
