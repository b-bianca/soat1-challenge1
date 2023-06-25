package usecases

import (
	"soat1-challenge1/internal/core/domain"
	"soat1-challenge1/internal/core/ports"
)

type clientUseCase struct {
	clientRepo ports.ClientRepository
}

func NewClientUseCase(clientRepo ports.ClientRepository) ports.ClientUseCase {
	return &clientUseCase{
		clientRepo: clientRepo,
	}
}

func (t *clientUseCase) Get(id string) (*domain.Client, error) {
	client, err := t.clientRepo.Get(id)
	if err != nil {
		//TODO: add logs
		return nil, err
	}

	return client, nil
}

func (t *clientUseCase) List() ([]domain.Client, error) {
	clients, err := t.clientRepo.List()
	if err != nil {
		//TODO: add logs
		return nil, err
	}

	return clients, nil
}

func (t *clientUseCase) Create(name, CPF, email string) (*domain.Client, error) {
	client := domain.NewClient("", name, CPF, email)

	_, err := t.clientRepo.Create(client)
	if err != nil {
		//TODO: add logs
		return nil, err
	}

	return client, nil
}
