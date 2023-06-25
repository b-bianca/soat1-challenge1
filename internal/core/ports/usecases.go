package ports

import "soat1-challenge1/internal/core/domain"

type ClientUseCase interface {
	Get(id string) (*domain.Client, error)
	List() ([]domain.Client, error)
	Create(name, CPF, email string) (*domain.Client, error)
}
