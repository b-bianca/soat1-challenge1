package domain

import (
	"soat1-challenge1/internal/handlers/models"
	"time"
)

// Order domain table model
type Order struct {
	ID         int                `gorm:"not null"`
	Status     models.OrderStatus `gorm:"not null"`
	OrderItems []*OrderItems      `gorm:"foreignKey:OrderID"`
	CustomerID int
	CreatedAt  time.Time `gorm:"not null;autoCreateTime"`
	UpdatedAt  time.Time `gorm:"not null;autoUpdateTime"`
}

type OrderItems struct {
	ID        int       `gorm:"not null"`
	OrderID   int       `gorm:"not null"`
	ProductID int       `gorm:"not null"`
	Quantity  int       `gorm:"not_null"`
	CreatedAt time.Time `gorm:"not null;autoCreateTime"`
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime"`
}

type Item struct {
	ProductID int `gorm:"not null"`
	Quantity  int `gorm:"not_null"`
}

// OrderResponseList summary list
type OrderResponseList struct {
	Result []*Order
	Count  int64
}
