package order

import (
	"soat1-challenge1/internal/core/ports"

	"github.com/gin-gonic/gin"
)

// Handler provides order funcionalities
type Handler struct {
	useCase ports.OrderUseCase
}

// NewHandler is the order handler builder
func NewHandler(u ports.OrderUseCase) *Handler {
	return &Handler{
		useCase: u,
	}
}

// RegisterRoutes register routes
func (h *Handler) RegisterRoutes(routes *gin.RouterGroup) {
	orderRoute := routes.Group("/order")
	orderRoute.GET("", h.List)
}
