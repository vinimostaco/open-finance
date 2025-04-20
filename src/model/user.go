package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
    ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
    Name      string    `gorm:"type:text;not null"`
    Email     string    `gorm:"type:text;unique;not null"`
    Password  string    `gorm:"type:text;not null"`
    CreatedAt time.Time
    UpdatedAt time.Time
}

type RegisterUserInput struct {
    Name     string `json:"nome"`
    Email    string `json:"email"`
    Password string `json:"senha"`
}

type LoginUserInput struct {
    Email    string `json:"email"`
    Password string `json:"senha"`
}