package product

import (
	"net/http"
	"soat1-challenge1/internal/core/domain"
	"soat1-challenge1/internal/handlers/dto"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateCategory(ctx *gin.Context) {
	var input *dto.CategoryRequestDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		return
	}

	domain := &domain.Category{
		Category: input.Category,
	}

	res, err := h.useCase.CreateCategory(ctx, domain)
	if err != nil {
		return
	}

	output := &dto.CategoryResponseDTO{
		ID:        res.ID,
		Category:  res.Category,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, output)
}
