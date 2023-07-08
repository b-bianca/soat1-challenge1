package domain

import "time"

type Customer struct {
	ID        int    `gorm:"not null"`
	Name      string `gorm:"not null"`
	CPF       string
	Email     string
	CreatedAt time.Time `gorm:"not null;autoCreateTime"`
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime"`
}
