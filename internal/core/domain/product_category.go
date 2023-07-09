package domain

import (
	"time"
)

// Category domain table model
type Category struct {
	ID        int       `gorm:"not null"`
	Name      string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null;autoCreateTime"`
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime"`
}

// CategoriesResponseList summary list
type CategoriesResponseList struct {
	Result []*Category
	Count  int64
}