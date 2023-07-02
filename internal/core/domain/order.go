package domain

import "time"

// Order domain table model
type Order struct {
	ID        int
	Confirmed bool
	Paid      bool
	StatusID  int
	ClientID  int
	CreatedAt time.Time
	UpdatedAt time.Time
}

// OrderResponseList summary list
type OrderResponseList struct {
	Result []*Order
	Count  int64
}
