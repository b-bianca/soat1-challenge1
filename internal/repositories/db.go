package repository

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewRepository creates a new gorm mysql db repository instance.
func NewRepository() (repo *Repository) {

	gdb, _ := gorm.Open(mysql.Open("mysql://user:user@tcp(127.0.0.1:3306)/restaurante"), &gorm.Config{})

	repo = New(gdb)

	return
}
