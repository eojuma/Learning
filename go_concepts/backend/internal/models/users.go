package models

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrEmailTaken       = errors.New("email address already exist")
	ErrUsernameTaken    = errors.New("username already exist")
	ErrPhoneNumberTaken = errors.New("Phone number already exist")
	ErrIDNumberTaken    = errors.New("ID Number already exist")
	ErrUserNotFound =errors.New("user does not exist")
)

type User struct {
	ID             uuid.UUID
	Name           string
	Username       string
	EmailAddress   string
	PhoneNumber    string
	IDNumber       string
	HashedPassword string
}
