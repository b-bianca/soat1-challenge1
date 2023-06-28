package domain

import "time"

// Order domain table model
type Order struct {
	ID        string `gorm:"not null, primaryKey"`
	Confirmed bool   `gorm:"not null"`
	Paid      bool   `gorm:"not null"`
	StatusID  string `gorm:"not null"`
	ClientID  int
	CreatedAt time.Time `gorm:"not null;default:now()"`
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime"`
}

// OrderResponseList summary list
type OrderResponseList struct {
	Result []*Order
	Count  int64
}
