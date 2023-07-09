package customer

import (
	"fmt"
	"net/http"
	"soat1-challenge1/internal/core/domain"
	"soat1-challenge1/internal/handlers/dto"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateCustomer(ctx *gin.Context) {
	var input *dto.CustomerRequestDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		return
	}

	domain := &domain.Customer{
		Name:  input.Name,
		Email: input.Email,
		CPF:   input.CPF,
	}

	res, err := h.useCase.CreateCustomer(ctx, domain)
	if err != nil {
		return
	}

	output := &dto.CustomerResponseDTO{
		ID:        res.ID,
		Name:      res.Name,
		Email:     res.Email,
		CPF:       res.CPF,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}

	ctx.JSON(http.StatusCreated, output)
}

func (h *Handler) RetrieveCustomer(ctx *gin.Context) {
	cpf := ctx.Param("cpf")
	fmt.Print(cpf)
	domain := &domain.Customer{
		CPF: cpf,
	}

	res, err := h.useCase.RetrieveCustomer(ctx, domain)
	if err != nil {
		return
	}

	output := &dto.CustomerResponseDTO{
		ID:        res.ID,
		Name:      res.Name,
		Email:     res.Email,
		CPF:       res.CPF,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, output)
}
