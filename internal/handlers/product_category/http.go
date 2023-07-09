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
		Name: input.Name,
	}

	res, err := h.useCase.CreateCategory(ctx, domain)
	if err != nil {
		return
	}

	output := &dto.CategoryResponseDTO{
		ID:        res.ID,
		Name:      res.Name,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, output)
}

func (h *Handler) GetCategories(ctx *gin.Context) {
	res, err := h.useCase.GetCategories(ctx)
	if err != nil {
		return
	}

	responseItems := make([]*dto.CategoryResponseDTO, 0, len(res.Result))

	for _, item := range res.Result {
		responseItems = append(responseItems, &dto.CategoryResponseDTO{
			ID:        item.ID,
			Name:      item.Name,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}

	output := &dto.CategoriesResponseList{
		Result: responseItems,
		Count:  res.Count,
	}

	ctx.JSON(http.StatusOK, output)
}
