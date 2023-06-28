package manage

import (
	"soat1-challenge1/internal/core/ports"
	o "soat1-challenge1/internal/handlers/order"

	"github.com/gin-gonic/gin"
)

type apps interface {
	RegisterRoutes(routes *gin.RouterGroup)
}

// Manage holds client instance
type Manage struct {
	order apps
}

// UseCases config struct
type UseCases struct {
	Order ports.OrderUseCase
}

// New creates a new client to create new routes
func New(uc *UseCases) *Manage {

	orderHandler := o.NewHandler(uc.Order)

	return &Manage{
		order: orderHandler,
	}
}

// RegisterRoutes register routes and set its permissions
func (m *Manage) RegisterRoutes(group *gin.RouterGroup) {
	m.order.RegisterRoutes(group)
}
