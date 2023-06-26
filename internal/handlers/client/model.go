package client

import "soat1-challenge1/internal/core/domain"

type Client struct {
	ID   string `json:"id"`
	Name string `json:"title"`
	CPF  string `json:"description"`
}

type ClientList []Client

func (td *Client) FromDomain(client *domain.Client) {
	if td == nil {
		td = &Client{}
	}

	td.ID = client.ID
	td.Name = client.Name
	td.CPF = client.CPF
}

func (td *Client) Clientmain() *domain.Client {
	if td == nil {
		td = &Client{}
	}

	return &domain.Client{
		ID:   td.ID,
		Name: td.Name,
		CPF:  td.CPF,
	}
}

func (td ClientList) FromDomain(tdms []domain.Client) ClientList {
	for _, t := range tdms {
		client := Client{}
		client.FromDomain(&t)
		td = append(td, client)
	}

	return td
}
