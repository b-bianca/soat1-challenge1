package dto

import (
	"soat1-challenge1/internal/core/domain"
	"soat1-challenge1/internal/handlers/models"
	"time"
)

// OrderResponseDTO is the struct responsible to marshal to json body
type OrderResponseDTO struct {
	ID         int                  `json:"id"`
	Status     models.OrderStatus   `json:"status"`
	CustomerID int                  `json:"customer_id"`
	Items      []*domain.OrderItems `json:"items"`
	CreatedAt  time.Time            `json:"created_at"`
	UpdatedAt  time.Time            `json:"updated_at"`
}

type OrderRequestDTO struct {
	CustomerID int        `json:"customer_id"`
	Items      []ItemsDTO `json:"items"`
}

type ItemsDTO struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

// OrderResponseList is the struct responsible to marshal to json the response
type OrderResponseList struct {
	Result []*OrderResponseDTO `json:"result"`
	Count  int64               `json:"count"`
}

type CategoriesResponseList struct {
	Result []*CategoryResponseDTO `json:"result"`
	Count  int64                  `json:"count"`
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

type ProductResponseList struct {
	Result []*ProductResponseDTO `json:"result"`
	Count  int64                 `json:"count"`
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

type PaymentDTO struct {
	OrderID int                  `json:"order_id"`
	Price   float64              `json:"price"`
	Status  models.PaymentStatus `json:"status"`
}
