package domain

import (
	"time"
)

// Product domain table model
type Product struct {
	ID          int       `gorm:"not null"`
	Name        string    `gorm:"not null"`
	Description string    `gorm:"not null"`
	CategoryID  string    `gorm:"not null"`
	Price       float64   `gorm:"not null"`
	CreatedAt   time.Time `gorm:"not null;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"not null;autoUpdateTime"`
}
