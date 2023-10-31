package order

import (
	"net/http"
	"soat1-challenge1/internal/handlers/dto"
	"soat1-challenge1/internal/handlers/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) MakePayment(ctx *gin.Context) {
	orderID := ctx.Param("id")
	id, _ := strconv.Atoi(orderID)

	output := &dto.PaymentDTO{
		OrderID: id,
		Status:  models.Approved,
	}

	err := h.useCase.UpdateOrderStatus(ctx, id, models.Approved)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, output)
}
