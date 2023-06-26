package client

import (
	"database/sql"
	"fmt"
	"soat1-challenge1/internal/core/domain"
	"soat1-challenge1/internal/core/ports"
)

type clientMysql struct {
	ID   string
	Name string
	CPF  string
}

type clientListMysql []clientMysql

func (m *clientMysql) Clientmain() *domain.Client {
	return &domain.Client{
		ID:   m.ID,
		Name: m.Name,
		CPF:  m.CPF,
	}
}
func (m *clientMysql) FromDomain(client *domain.Client) {
	if m == nil {
		m = &clientMysql{}
	}

	m.ID = client.ID
	m.Name = client.Name
	m.CPF = client.CPF
}

func (m clientListMysql) Clientmain() []domain.Client {
	clients := make([]domain.Client, len(m))
	for k, td := range m {
		client := td.Clientmain()
		clients[k] = *client
	}

	return clients
}

type clientMysqlRepo struct {
	db *sql.DB
}

func NewClientMysqlRepo(db *sql.DB) ports.ClientRepository {
	return &clientMysqlRepo{
		db: db,
	}
}

func (m *clientMysqlRepo) Get(id string) (*domain.Client, error) {
	var client clientMysql = clientMysql{}
	sqsS := fmt.Sprintf("SELECT id, title, description FROM client WHERE id = '%s'", id)

	result := m.db.QueryRow(sqsS)
	if result.Err() != nil {
		return nil, result.Err()
	}

	if err := result.Scan(&client.ID, &client.Name, &client.CPF); err != nil {
		return nil, err
	}

	return client.Clientmain(), nil
}

func (m *clientMysqlRepo) List() ([]domain.Client, error) {
	var clients clientListMysql
	sqsS := "SELECT id, title, description FROM client"

	result, err := m.db.Query(sqsS)
	if err != nil {
		return nil, err
	}

	if result.Err() != nil {
		return nil, result.Err()
	}

	for result.Next() {
		client := clientMysql{}

		if err := result.Scan(&client.ID, &client.Name, &client.CPF); err != nil {
			return nil, err
		}

		clients = append(clients, client)
	}

	return clients.Clientmain(), nil
}

func (m *clientMysqlRepo) Create(client *domain.Client) (*domain.Client, error) {
	sqlS := "INSERT INTO client (id, title, description) VALUES (?, ?, ?)"

	_, err := m.db.Exec(sqlS, client.ID, client.Name, client.CPF)

	if err != nil {
		return nil, err
	}

	return client, nil
}
