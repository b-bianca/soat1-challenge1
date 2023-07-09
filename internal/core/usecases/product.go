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

// Create create and persist product data
func (u *useCaseProduct) Create(ctx context.Context, input *domain.Product) (*domain.Product, error) {
	p, err := u.repository.Create(ctx, input)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// Update updates a product and persist new data
func (u *useCaseProduct) Update(ctx context.Context, input *domain.Product) error {
	if err := u.repository.Update(ctx, input); err != nil {
		return err
	}

	return nil
}

// Delete remove a product data
func (u *useCaseProduct) Delete(ctx context.Context, input *domain.Product) error {
	if err := u.repository.Delete(ctx, input); err != nil {
		return err
	}

	return nil
}
