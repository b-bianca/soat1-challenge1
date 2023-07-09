package manage

import (
	"soat1-challenge1/internal/core/ports"
	c "soat1-challenge1/internal/handlers/customer"
	o "soat1-challenge1/internal/handlers/order"
	p "soat1-challenge1/internal/handlers/product"
	pc "soat1-challenge1/internal/handlers/product_category"

	"github.com/gin-gonic/gin"
)

type apps interface {
	RegisterRoutes(routes *gin.RouterGroup)
}

// Manage holds client instance
type Manage struct {
	order    apps
	product  apps
	productCategory  apps
	customer apps
}

// UseCases config struct
type UseCases struct {
	Order    ports.OrderUseCase
	Product  ports.ProductUseCase
	ProductCategory  ports.ProductCategoryUseCase
	Customer ports.CustomerUseCase
}

// New creates a new client to create new routes
func New(uc *UseCases) *Manage {

	orderHandler := o.NewHandler(uc.Order)
	productHandler := p.NewHandler(uc.Product)
	customerHandler := c.NewHandler(uc.Customer)
	productCategoryHandler := pc.NewHandler(uc.ProductCategory)

	return &Manage{
		order:    orderHandler,
		product:  productHandler,
		customer: customerHandler,
		productCategory: productCategoryHandler,
	}
}

// RegisterRoutes register routes and set its permissions
func (m *Manage) RegisterRoutes(group *gin.RouterGroup) {
	m.order.RegisterRoutes(group)
	m.product.RegisterRoutes(group)
	m.customer.RegisterRoutes(group)
	m.productCategory.RegisterRoutes(group)
}
