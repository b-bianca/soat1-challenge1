package ports

import (
	"context"
	"soat1-challenge1/internal/core/domain"
)

type CustomerUseCase interface {
	CreateCustomer(ctx context.Context, c *domain.Customer) (*domain.Customer, error)
	RetrieveCustomer(ctx context.Context, c *domain.Customer) (*domain.Customer, error)
}

// OrderUseCase is the interface for order repository
type OrderUseCase interface {
	CreateOrder(ctx context.Context, order *domain.Order) (*domain.Order, error)
	CreateOrderItens(ctx context.Context, order *domain.OrderItens) (*domain.OrderItens, error)
	List(context.Context) (*domain.OrderResponseList, error)
}

// ProductUseCase is the interface for product repository
type ProductUseCase interface {
	Create(ctx context.Context, product *domain.Product) (*domain.Product, error)
	Update(ctx context.Context, product *domain.Product) error
	Delete(ctx context.Context, pdt *domain.Product) error
	GetProducts(context.Context) (*domain.ProductResponseList, error)
}

// ProductCategoryUseCase is the interface for product_category repository
type ProductCategoryUseCase interface {
	CreateCategory(ctx context.Context, category *domain.Category) (*domain.Category, error)
	GetCategories(context.Context) (*domain.CategoriesResponseList, error)
}
