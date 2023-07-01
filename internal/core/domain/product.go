package domain

import (
	"time"

	"github.com/google/uuid"
)

// Product domain table model
type Product struct {
	ID          uuid.UUID
	Name        string
	Description string
	CategoryID  string
	Price       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Category domain table model
type Category struct {
	ID        int
	Category  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
