package product

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

	ctx.JSON(http.StatusOK, output)
}

func (h *Handler) Create(ctx *gin.Context) {
	var input *dto.ProductRequestDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		return
	}

	domain := &domain.Product{
		Name:        input.Name,
		Description: input.Description,
		CategoryID:  input.CategoryID,
		Price:       input.Price,
	}

	res, err := h.useCase.Create(ctx, domain)
	if err != nil {
		return
	}

	output := &dto.ProductResponseDTO{
		ID:          res.ID,
		Name:        res.Name,
		Description: res.Description,
		CategoryID:  res.CategoryID,
		Price:       res.Price,
		CreatedAt:   res.CreatedAt,
		UpdatedAt:   res.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, output)
}

func (h *Handler) Update(ctx *gin.Context) {
	var input *dto.ProductRequestDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		return
	}
	idParam := ctx.Param("id")
	id, _ := strconv.Atoi(idParam)

	domain := &domain.Product{
		ID:          id,
		Name:        input.Name,
		Description: input.Description,
		CategoryID:  input.CategoryID,
		Price:       input.Price,
	}

	err := h.useCase.Update(ctx, domain)
	if err != nil {
		return
	}

	ctx.JSON(http.StatusNoContent, "")
}

func (h *Handler) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, _ := strconv.Atoi(idParam)

	domain := &domain.Product{
		ID: id,
	}

	err := h.useCase.Delete(ctx, domain)
	if err != nil {
		return
	}

	ctx.JSON(http.StatusNoContent, "")
}
