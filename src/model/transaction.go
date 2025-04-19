package model

import (
	"time"

	"github.com/google/uuid"
)


type Transaction struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Title     string    `gorm:"type:text;not null"`
	Amount    float64   `gorm:"type:numeric(10,2);not null"`
	Type      string    `gorm:"type:text;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}	

type AddTransactionInput struct {
    Nome  string  `json:"nome"`
    Valor float64 `json:"valor"`
    Tipo  string  `json:"tipo"`
}

type GetTransactionByNameInput struct {
	Nome string `json:"nome"`
}

type GetTransactionByTypeInput struct {
	Tipo string `json:"tipo"`
}