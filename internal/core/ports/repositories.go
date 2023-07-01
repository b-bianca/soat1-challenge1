package ports

import (
	"context"
	"soat1-challenge1/internal/core/domain"
)

// OrderRepository is the interface for order database
type OrderRepository interface {
	List(context.Context) (*domain.OrderResponseList, error)
}

// ProductRepository is the interface for product database
type ProductRepository interface {
	Create(ctx context.Context, product *domain.Product) (*domain.Product, error)
	CreateCategory(ctx context.Context, category *domain.Category) (*domain.Category, error)
}
