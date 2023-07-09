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
	CreateOrderItems(ctx context.Context, order []*domain.OrderItems) ([]*domain.OrderItems, error)
}

// ProductRepository is the interface for product database
type ProductRepository interface {
	Create(ctx context.Context, product *domain.Product) (*domain.Product, error)
	Update(ctx context.Context, product *domain.Product) error
	Delete(ctx context.Context, pdt *domain.Product) error
	GetProducts(context.Context) (*domain.ProductResponseList, error)
}

// ProductCategoryRepository is the interface for product_category database
type ProductCategoryRepository interface {
	CreateCategory(ctx context.Context, category *domain.Category) (*domain.Category, error)
	GetCategories(context.Context) (*domain.CategoriesResponseList, error)
	GetCategoryProducts(ctx context.Context, category *domain.Category) (*domain.ProductResponseList, error)
}
