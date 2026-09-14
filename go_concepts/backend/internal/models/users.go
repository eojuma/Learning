package models

import "github.com/google/uuid"

type User struct {
	ID             uuid.UUID
	Name           string
	Username       string
	EmailAddress   string
	PhoneNumber    string
	IDNumber       string
	HashedPassword string
}
