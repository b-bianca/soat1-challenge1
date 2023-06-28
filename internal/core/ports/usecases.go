package ports

import (
	"context"
	"soat1-challenge1/internal/core/domain"
)

type ClientUseCase interface {
	Get(id string) (*domain.Client, error)
	List() ([]domain.Client, error)
	Create(name, CPF, email string) (*domain.Client, error)
}

// OrderUseCase is the interface for order repository
type OrderUseCase interface {
	List(context.Context) (*domain.OrderResponseList, error)
}
