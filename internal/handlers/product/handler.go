package product

import (
	"soat1-challenge1/internal/core/ports"

	"github.com/gin-gonic/gin"
)

// Handler provides product funcionalities
type Handler struct {
	useCase ports.ProductUseCase
}

// NewHandler is the product handler builder
func NewHandler(u ports.ProductUseCase) *Handler {
	return &Handler{
		useCase: u,
	}
}

// RegisterRoutes register routes
func (h *Handler) RegisterRoutes(routes *gin.RouterGroup) {
	productRoute := routes.Group("/products")
	productRoute.GET("", h.GetProducts)
	productRoute.POST("", h.Create)
	productRoute.PUT("/:id", h.Update)
	productRoute.DELETE("/:id", h.Delete)
}
