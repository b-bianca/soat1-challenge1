package product

import (
	"net/http"
	"soat1-challenge1/internal/core/domain"
	"soat1-challenge1/internal/handlers/dto"
	"strconv"

	"github.com/gin-gonic/gin"
)

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

	ctx.JSON(http.StatusCreated, output)
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

	ctx.JSON(http.StatusOK, "")
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

	ctx.JSON(http.StatusOK, "")
}

func (h *Handler) GetProducts(ctx *gin.Context) {
	res, err := h.useCase.GetProducts(ctx)
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
