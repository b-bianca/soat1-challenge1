package order

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"soat1-challenge1/internal/core/domain"
	"soat1-challenge1/internal/handlers/dto"
	"soat1-challenge1/internal/handlers/models"
	"strconv"

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
		orderItems, err := h.useCase.GetOrderItems(ctx, item.ID)
		if err != nil {
			return
		}

		responseItems = append(responseItems, &dto.OrderResponseDTO{
			ID:         item.ID,
			Status:     item.Status,
			Items:      orderItems,
			CustomerID: item.CustomerID,
			CreatedAt:  item.CreatedAt,
			UpdatedAt:  item.UpdatedAt,
		})
		fmt.Printf("%+v", responseItems)
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
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	domainOrder := &domain.Order{
		CustomerID: input.CustomerID,
		Status:     models.Pending,
	}

	res, err := h.useCase.CreateOrder(ctx, domainOrder)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	fmt.Print(res)

	orderItems := make([]*domain.OrderItems, 0, MaxQuantityPerOrder)

	for _, item := range input.Items {
		fmt.Print(item)
		orderItems = append(orderItems, &domain.OrderItems{
			OrderID:   res.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		})
	}
	fmt.Print(orderItems)
	resItems, err := h.useCase.CreateOrderItems(ctx, orderItems)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.AddParam("id", strconv.Itoa(res.ID))
	url := fmt.Sprintf("http://localhost:8080/api/v1/orders/%s/payments", strconv.Itoa(res.ID))
	respPayment, err := http.Post(url, "application/json", nil)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	var payment *dto.PaymentDTO
	dec := json.NewDecoder(respPayment.Body)
	if err := dec.Decode(&payment); err != nil && err != io.EOF {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	switch payment.Status {
	case models.Approved:
		res.Status = models.Received
	case models.Rejected:
		res.Status = models.Cancelled
	}

	err = h.useCase.UpdateOrderStatus(ctx, res.ID, res.Status.String())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	output := &dto.OrderResponseDTO{
		ID:         res.ID,
		Status:     res.Status,
		CustomerID: res.CustomerID,
		CreatedAt:  res.CreatedAt,
		UpdatedAt:  res.UpdatedAt,
		Items:      resItems,
	}

	ctx.JSON(http.StatusCreated, output)
}

func (h *Handler) MakePayment(ctx *gin.Context) {
	orderID := ctx.Param("id")
	id, _ := strconv.Atoi(orderID)

	output := &dto.PaymentDTO{
		OrderID: id,
		Status:  models.Approved,
	}

	ctx.JSON(http.StatusOK, output)
}
