package repository

import (
	"context"
	"soat1-challenge1/internal/core/domain"

	"gorm.io/gorm"
)

// Order struct type
type Order struct {
	db *gorm.DB
}

// NewOrderRepository instantiates order repository
func NewOrderRepository(db *gorm.DB) *Order {
	return &Order{db}
}

// List retrives all orders.
func (o *Order) List(ctx context.Context) (*domain.OrderResponseList, error) {
	dbFn := o.db.WithContext(ctx)

	var count int64
	var order []*domain.Order

	result := dbFn.Table("order").Find(&order).Count(&count)

	if result.Error != nil {
		return nil, result.Error
	}
	return &domain.OrderResponseList{
		Result: order,
		Count:  count,
	}, nil
}
