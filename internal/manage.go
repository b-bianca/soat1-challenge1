package manage

import (
	"soat1-challenge1/internal/core/ports"
	c "soat1-challenge1/internal/handlers/customer"
	o "soat1-challenge1/internal/handlers/order"
	p "soat1-challenge1/internal/handlers/product"

	"github.com/gin-gonic/gin"
)

type apps interface {
	RegisterRoutes(routes *gin.RouterGroup)
}

// Manage holds client instance
type Manage struct {
	order    apps
	product  apps
	customer apps
}

// UseCases config struct
type UseCases struct {
	Order    ports.OrderUseCase
	Product  ports.ProductUseCase
	Customer ports.CustomerUseCase
}

// New creates a new client to create new routes
func New(uc *UseCases) *Manage {

	orderHandler := o.NewHandler(uc.Order)
	productHandler := p.NewHandler(uc.Product)
	customerHandler := c.NewHandler(uc.Customer)

	return &Manage{
		order:    orderHandler,
		product:  productHandler,
		customer: customerHandler,
	}
}

// RegisterRoutes register routes and set its permissions
func (m *Manage) RegisterRoutes(group *gin.RouterGroup) {
	m.order.RegisterRoutes(group)
	m.product.RegisterRoutes(group)
	m.customer.RegisterRoutes(group)
}
