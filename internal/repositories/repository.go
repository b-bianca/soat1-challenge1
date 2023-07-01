package repository

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

// Repository is a service to abstract mysql layer
type Repository struct {
	Conn    *gorm.DB
	connSQL *sql.DB
	Order   *Order
	Product *Product
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
		Product: NewProductRepository(gdb),
	}
}

// SetMaxOpenConns sets the maximum number of open connections to the database.
func (r *Repository) SetMaxOpenConns(n int) {
	r.connSQL.SetMaxOpenConns(n)
}

// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
func (r *Repository) SetMaxIdleConns(n int) {
	r.connSQL.SetMaxIdleConns(n)
}

// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
func (r *Repository) SetConnMaxLifetime(t time.Duration) {
	r.connSQL.SetConnMaxLifetime(t)
}

// SetConnMaxIdleTime sets the maximum amount of time a connection may be idle.
func (r *Repository) SetConnMaxIdleTime(t time.Duration) {
	r.connSQL.SetConnMaxIdleTime(t)
}
