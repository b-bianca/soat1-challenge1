package repository

import (
	"database/sql"

	"gorm.io/gorm"
)

// Repository is a service to abstract mysql layer
type Repository struct {
	Conn    *gorm.DB
	connSQL *sql.DB
	Order   *Order
}

// New creates a new repository
func New(gdb *gorm.DB) *Repository {
	conn, err := gdb.DB()
	if err != nil {
		return nil
	}

	return &Repository{
		Conn:    gdb,
		connSQL: conn,
		Order:   NewOrderRepository(gdb),
	}
}
