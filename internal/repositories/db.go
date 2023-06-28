package repository

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	MaxOpenConns       = 2
	MaxIdleConns       = 2
	ConnMaxLifetimeSec = 100
	ConnMaxIdleSec     = 100
)

// NewRepository creates a new gorm mysql db repository instance.
func NewRepository() (repo *Repository) {

	gdb, err := gorm.Open(mysql.Open("user123:userpass@tcp(127.0.0.1:3306)/restaurant"), &gorm.Config{})
	if err != nil {
		log.Panicf("failed to create mysql db: %v", err)
	}

	repo = New(gdb)

	if MaxOpenConns > 0 {
		repo.SetMaxOpenConns(MaxOpenConns)
	}

	if MaxIdleConns > 0 {
		repo.SetMaxIdleConns(MaxIdleConns)
	}

	if ConnMaxLifetimeSec > 0 {
		repo.SetConnMaxLifetime(ConnMaxLifetimeSec)
	}

	if ConnMaxIdleSec > 0 {
		repo.SetConnMaxIdleTime(ConnMaxIdleSec)
	}

	return
}
