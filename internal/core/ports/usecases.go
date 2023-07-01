package ports

import (
	"context"
	"soat1-challenge1/internal/core/domain"
)

type ClientUseCase interface {
	Get(id string) (*domain.Client, error)
	List() ([]domain.Client, error)
	Create(name, CPF, email string) (*domain.Client, error)
}

// OrderUseCase is the interface for order repository
type OrderUseCase interface {
	List(context.Context) (*domain.OrderResponseList, error)
}

// ProductUseCase is the interface for product repository
type ProductUseCase interface {
	Create(ctx context.Context, product *domain.Product) (*domain.Product, error)
	CreateCategory(ctx context.Context, category *domain.Category) (*domain.Category, error)
}
