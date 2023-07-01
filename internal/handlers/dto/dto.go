package dto

import (
	"time"

	"github.com/google/uuid"
)

// OrderResponseDTO is the struct responsible to marshal to json body
type OrderResponseDTO struct {
	ID        uuid.UUID `json:"id"`
	Confirmed bool      `json:"confirmed"`
	Paid      bool      `json:"paid"`
	StatusID  string    `json:"status_id"`
	ClientID  string    `json:"client"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// OrderResponseList is the struct responsible to marshal to json the response
type OrderResponseList struct {
	Result []*OrderResponseDTO `json:"result"`
	Count  int64               `json:"count"`
}

type CategoryRequestDTO struct {
	Category string `json:"category"`
}

type CategoryResponseDTO struct {
	ID        int       `json:"id"`
	Category  string    `json:"category"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
