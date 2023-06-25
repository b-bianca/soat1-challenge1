package domain

import (
	"fmt"
)

type Client struct {
	ID          string
	Name        string
	CPF         string
	Email       string
	DateCreated string
}

func NewClient(id string, name, CPF, email string) *Client {
	return &Client{
		ID:          id,
		Name:        name,
		CPF:         CPF,
		Email:       email,
		DateCreated: "",
	}
}

func (c *Client) String() string {
	return fmt.Sprintf("%s - %s", c.Name, c.CPF)
}
