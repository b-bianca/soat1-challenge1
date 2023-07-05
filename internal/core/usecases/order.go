package usecases

import (
	"context"
	"soat1-challenge1/internal/core/domain"
	"soat1-challenge1/internal/core/ports"
)

type useCaseOrder struct {
	repository ports.OrderRepository
}

// NewOrderUseCase is responsible for all use cases for orders
func NewOrderUseCase(orderRepo ports.OrderRepository) ports.OrderUseCase {
	return &useCaseOrder{
		repository: orderRepo,
	}
}

// List retrieves all orders
func (u *useCaseOrder) List(ctx context.Context) (*domain.OrderResponseList, error) {
	res, err := u.repository.List(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}
