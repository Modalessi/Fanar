package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID       *uuid.UUID
	Name     string
	Email    string
	Password string
}

func NewUser(name string, email string, password string) *User {
	return &User{
		Name:     name,
		Email:    email,
		Password: password,
	}
}
