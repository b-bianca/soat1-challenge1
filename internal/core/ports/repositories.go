package ports

import "soat1-challenge1/internal/core/domain"

type ClientRepository interface {
	Get(id string) (*domain.Client, error)
	List() ([]domain.Client, error)
	Create(client *domain.Client) (*domain.Client, error)
}
