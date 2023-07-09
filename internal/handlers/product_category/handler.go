package product

import (
	"soat1-challenge1/internal/core/ports"

	"github.com/gin-gonic/gin"
)

// Handler provides product funcionalities
type Handler struct {
	useCase ports.ProductCategoryUseCase
}

// NewHandler is the product handler builder
func NewHandler(u ports.ProductCategoryUseCase) *Handler {
	return &Handler{
		useCase: u,
	}
}

// RegisterRoutes register routes
func (h *Handler) RegisterRoutes(routes *gin.RouterGroup) {
	productCategoryRoute := routes.Group("/categories")
	productCategoryRoute.POST("", h.CreateCategory)
	productCategoryRoute.GET("", h.GetCategories)
}
