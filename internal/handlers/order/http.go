package order

import (
	"net/http"
	"soat1-challenge1/internal/core/domain"
	"soat1-challenge1/internal/handlers/dto"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	MaxQuantityPerOrder = 10
)

func (h *Handler) List(ctx *gin.Context) {
	res, err := h.useCase.List(ctx)
	if err != nil {
		return
	}

	responseItems := make([]*dto.OrderResponseDTO, 0, len(res.Result))

	for _, item := range res.Result {
		responseItems = append(responseItems, &dto.OrderResponseDTO{
			ID:         item.ID,
			Status:     item.Status,
			CustomerID: item.CustomerID,
			CreatedAt:  item.CreatedAt,
			UpdatedAt:  item.UpdatedAt,
		})
	}

	output := &dto.OrderResponseList{
		Result: responseItems,
		Count:  res.Count,
	}

	ctx.JSON(http.StatusOK, output)
}

func (h *Handler) CreateOrder(ctx *gin.Context) {
	var input *dto.OrderRequestDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		return
	}

	domainOrder := &domain.Order{
		CustomerID: input.CustomerID,
	}

	res, err := h.useCase.CreateOrder(ctx, domainOrder)
	if err != nil {
		return
	}

	responseItems := make([]*domain.Itens, 0, MaxQuantityPerOrder)

	for _, item := range input.Itens {
		responseItems = append(responseItems, &domain.Itens{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		})
	}

	domainOrderItem := &domain.OrderItens{
		OrderID: res.ID,
		Itens:   responseItems,
	}

	itens, err := h.useCase.CreateOrderItens(ctx, domainOrderItem)
	if err != nil {
		return
	}

	output := &dto.OrderResponseDTO{}

	ctx.JSON(http.StatusOK, output)
}
