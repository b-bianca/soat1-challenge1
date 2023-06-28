package order

import (
	"net/http"
	"soat1-challenge1/internal/handlers/dto"

	"github.com/gin-gonic/gin"
)

func (h *Handler) List(ctx *gin.Context) {
	res, err := h.useCase.List(ctx)
	if err != nil {
		return
	}

	responseItems := make([]*dto.OrderResponseDTO, 0, len(res.Result))

	for _, item := range res.Result {
		responseItems = append(responseItems, &dto.OrderResponseDTO{
			ID:        item.ID,
			Confirmed: item.Confirmed,
			Paid:      item.Paid,
			StatusID:  item.StatusID,
			ClientID:  item.ClientID,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}

	output := &dto.OrderResponseList{
		Result: responseItems,
		Count:  res.Count,
	}

	ctx.JSON(http.StatusOK, output)
}
