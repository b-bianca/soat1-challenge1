package repository

import (
	"context"
	"soat1-challenge1/internal/core/domain"

	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func (o *Order) GetOrderItems(ctx context.Context, orderID int) ([]*domain.OrderItems, error) {
	dbFn := o.db.WithContext(ctx)

	var orderItems []*domain.OrderItems

	result := dbFn.Table("order_item").Find(&orderItems).Where("order_id = ?", orderID)

	fmt.Sprintf("%+v", orderItems)
	if result.Error != nil {
		return nil, result.Error
	}

	return orderItems, nil
}

func (o *Order) CreateOrder(ctx context.Context, order *domain.Order) (*domain.Order, error) {
	dbFn := o.db.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true})

	if result := dbFn.Table("order").Create(order); result.Error != nil {
		return nil, result.Error
	}

	return order, nil
}

func (o *Order) CreateOrderItems(ctx context.Context, orderItems []*domain.OrderItems) ([]*domain.OrderItems, error) {
	dbFn := o.db.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true})

	if result := dbFn.Table("order_item").Create(orderItems); result.Error != nil {
		return nil, result.Error
	}

	return orderItems, nil
}

func (o *Order) UpdateOrderStatus(ctx context.Context, orderID int, status string) error {
	dbFn := o.db.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true})

	if result := dbFn.Table("order").Where("id = ?", orderID).Update("status", status); result.Error != nil {
		return result.Error
	}

	return nil
}
