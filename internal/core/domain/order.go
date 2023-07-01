package domain

import (
	"time"

	"github.com/google/uuid"
)

// Order domain table model
type Order struct {
	ID        uuid.UUID
	Confirmed bool
	Paid      bool
	StatusID  string
	ClientID  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// OrderResponseList summary list
type OrderResponseList struct {
	Result []*Order
	Count  int64
}
