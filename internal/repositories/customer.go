package repository

import (
	"context"
	"soat1-challenge1/internal/core/domain"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Customer struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *Customer {
	return &Customer{db}
}

func (p *Customer) CreateCustomer(ctx context.Context, c *domain.Customer) (*domain.Customer, error) {
	dbFn := p.db.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true})

	if result := dbFn.Table("customer").Create(c); result.Error != nil {
		return nil, result.Error
	}

	return c, nil
}
