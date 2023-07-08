package customer

import (
	"soat1-challenge1/internal/core/ports"

	"github.com/gin-gonic/gin"
)

// Handler provides product funcionalities
type Handler struct {
	useCase ports.CustomerUseCase
}

// NewHandler is the product handler builder
func NewHandler(u ports.CustomerUseCase) *Handler {
	return &Handler{
		useCase: u,
	}
}

// RegisterRoutes register routes
func (h *Handler) RegisterRoutes(routes *gin.RouterGroup) {
	customerRoute := routes.Group("/customer")
	customerRoute.POST("/", h.CreateCustomer)
	customerRoute.GET("/:cpf", h.RetrieveCustomer)
}
