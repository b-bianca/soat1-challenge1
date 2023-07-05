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
	Name string `json:"name"`
}

type CategoryResponseDTO struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ProductRequestDTO struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	CategoryID  string  `json:"category_id"`
	Price       float64 `json:"price"`
}

type ProductResponseDTO struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CategoryID  string    `json:"category_id"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CustomerRequestDTO struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	CPF   string `json:"cpf"`
}

type CustomerResponseDTO struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CPF       string    `json:"cpf"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
