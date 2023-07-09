package usecases

import (
	"context"
	"soat1-challenge1/internal/core/domain"
	"soat1-challenge1/internal/core/ports"
)

type useCaseProductCategory struct {
	repository ports.ProductCategoryRepository
}

// NewOrderUseCase is responsible for all use cases for orders
func NewProductCategoryUseCase(p ports.ProductCategoryRepository) ports.ProductCategoryUseCase {
	return &useCaseProductCategory{
		repository: p,
	}
}

// CreateCategory persist category data
func (u *useCaseProductCategory) CreateCategory(ctx context.Context, input *domain.Category) (*domain.Category, error) {
	res, err := u.repository.CreateCategory(ctx, input)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// List retrieves all categories
func (u *useCaseProductCategory) GetCategories(ctx context.Context) (*domain.CategoriesResponseList, error) {
	res, err := u.repository.GetCategories(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}