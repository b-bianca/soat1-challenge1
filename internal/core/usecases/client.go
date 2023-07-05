package usecases

import (
	"context"
	"soat1-challenge1/internal/core/domain"
	"soat1-challenge1/internal/core/ports"
)

type useCaseCustomer struct {
	repository ports.CustomerRepository
}

func NewCustomerUseCase(c ports.CustomerRepository) ports.CustomerUseCase {
	return &useCaseCustomer{
		repository: c,
	}
}

func (u *useCaseCustomer) CreateCustomer(ctx context.Context, c *domain.Customer) (*domain.Customer, error) {
	res, err := u.repository.CreateCustomer(ctx, c)
	if err != nil {
		return nil, err
	}

	return res, nil
}
