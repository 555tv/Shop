package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

// MongoDB
type User struct {
	ID           bson.ObjectID `bson:"_id,omitempty" json:"id"`
	FirstName    string        `bson:"firstName" json:"firstName"`
	LastName     string        `bson:"lastName" json:"lastName"`
	Email        string        `bson:"email" json:"email"`
	BirthDate    time.Time     `bson:"birthDate" json:"birthDate"`
	PasswordHash string        `bson:"passwordHash" json:"-"`
	CreatedAt    time.Time     `bson:"createdAt" json:"createdAt"`
	Role         string        `bson:"role"`
}

// Запросы
type UserRequest struct {
	FirstName string    `json:"firstName" validate:"required,min=2,max=30,alpha"`
	LastName  string    `json:"lastName" validate:"required,min=2,max=30,alpha"`
	BirthDate time.Time `json:"birthDate" validate:"required,birthdate"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"password" validate:"required,min=6"`
}

type UpdateUserRequest struct {
	FirstName string    `json:"firstName" validate:"required,min=2,max=30,alpha"`
	LastName  string    `json:"lastName" validate:"required,min=2,max=30,alpha"`
	BirthDate time.Time `json:"birthDate" validate:"required,birthdate"`
}

type SearchUserRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// Ответ
type UserResponse struct {
	ID        string    `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	BirthDate time.Time `json:"birthDate"`
	CreatedAt time.Time `json:"createdAt"`
	Email     string    `json:"email"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
