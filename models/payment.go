package models

import (
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	ID          uuid.UUID  `gorm:"type:uuid;primary_key;" json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `sql:"index" json:"deleted_at,omitempty"`
	CardNumber  string     `json:"card_number"`
	ExpiryMonth string     `json:"expiry_month"`
	ExpiryYear  string     `json:"expiry_year"`
	Amount      float64    `json:"amount"`
	Currency    string     `json:"currency"`
	Status      string     `json:"status"`
}

type ProcessPaymentRequest struct {
	CardNumber  string  `json:"card_number"`
	ExpiryMonth string  `json:"expiry_month"`
	ExpiryYear  string  `json:"expiry_year"`
	CVV         string  `json:"cvv"`
	Amount      float64 `json:"amount"`
	Currency    string  `json:"currency"`
}

type BankResponse struct {
	Status string `json:"status"`
}
