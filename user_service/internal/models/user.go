package models

import "time"

type User struct {
	ID           int
	FirstName    string
	LastName     string
	Email        string
	BirthDate    *time.Time
	PasswordHash string
	Role         string
	CreatedAt    time.Time
}
