package usecases

import (
	"context"
	"soat1-challenge1/internal/core/domain"
	"soat1-challenge1/internal/core/ports"
)

type useCaseProduct struct {
	repository ports.ProductRepository
}

// NewOrderUseCase is responsible for all use cases for orders
func NewProductUseCase(p ports.ProductRepository) ports.ProductUseCase {
	return &useCaseProduct{
		repository: p,
	}
}

// CreateCategory persist category data
func (u *useCaseProduct) CreateCategory(ctx context.Context, input *domain.Category) (*domain.Category, error) {
	res, err := u.repository.CreateCategory(ctx, input)
	if err != nil {
		//TODO: ADD LOGS
		return nil, err
	}

	return res, nil
}

// List retrieves all orders
func (u *useCaseProduct) Create(ctx context.Context, input *domain.Product) (*domain.Product, error) {
	return nil, nil
}
