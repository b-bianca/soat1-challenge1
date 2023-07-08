package ports

import (
	"context"
	"soat1-challenge1/internal/core/domain"
)

type CustomerRepository interface {
	CreateCustomer(ctx context.Context, c *domain.Customer) (*domain.Customer, error)
	RetrieveCustomer(ctx context.Context, c *domain.Customer) (*domain.Customer, error)
}

// OrderRepository is the interface for order database
type OrderRepository interface {
	List(context.Context) (*domain.OrderResponseList, error)
	CreateOrder(context.Context, *domain.Order) (*domain.Order, error)
	CreateOrderItens(ctx context.Context, order *domain.OrderItens) (*domain.OrderItens, error)
}

// ProductRepository is the interface for product database
type ProductRepository interface {
	CreateCategory(ctx context.Context, category *domain.Category) (*domain.Category, error)
	Create(ctx context.Context, product *domain.Product) (*domain.Product, error)
	Update(ctx context.Context, product *domain.Product) error
	Delete(ctx context.Context, pdt *domain.Product) error
}
