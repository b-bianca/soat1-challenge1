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

func (c *Customer) CreateCustomer(ctx context.Context, ctm *domain.Customer) (*domain.Customer, error) {
	dbFn := c.db.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true})

	if result := dbFn.Table("customer").Create(ctm); result.Error != nil {
		return nil, result.Error
	}

	return ctm, nil
}

func (c *Customer) RetrieveCustomer(ctx context.Context, ctm *domain.Customer) (*domain.Customer, error) {
	dbFn := c.db.WithContext(ctx)
	cpf := ctm.CPF

	if result := dbFn.Table("customer").Where("cpf = ?", cpf).Find(&ctm); result.Error != nil {
		return nil, result.Error
	}

	return ctm, nil
}
