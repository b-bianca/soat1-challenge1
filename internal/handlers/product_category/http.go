package product_category

import (
	"net/http"
	"soat1-challenge1/internal/core/domain"
	"soat1-challenge1/internal/handlers/dto"
	"strconv"

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

	ctx.JSON(http.StatusCreated, output)
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

func (h *Handler) GetCategoryProducts(ctx *gin.Context) {
	categoryID := ctx.Param("id")
	id, _ := strconv.Atoi(categoryID)
	domain := &domain.Category{
		ID: id,
	}
	res, err := h.useCase.GetCategoryProducts(ctx, domain)
	if err != nil {
		return
	}

	responseItems := make([]*dto.ProductResponseDTO, 0, len(res.Result))

	for _, item := range res.Result {
		responseItems = append(responseItems, &dto.ProductResponseDTO{
			ID:          item.ID,
			Name:        item.Name,
			Description: item.Description,
			CategoryID:  item.CategoryID,
			Price:       item.Price,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
		})
	}

	output := &dto.ProductResponseList{
		Result: responseItems,
		Count:  res.Count,
	}

	ctx.JSON(http.StatusOK, output)
}
