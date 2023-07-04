package dto

import (
	"time"
)

// OrderResponseDTO is the struct responsible to marshal to json body
type OrderResponseDTO struct {
	ID        int       `json:"id"`
	Confirmed bool      `json:"confirmed"`
	Paid      bool      `json:"paid"`
	StatusID  int       `json:"status_id"`
	ClientID  int       `json:"client"`
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
