package domain

import "time"

// Order domain table model
type Order struct {
	ID         int    `gorm:"not null"`
	Status     string `gorm:"type:enum('aguardando pagamento', 'pago');default:'aguardando pagamento'"`
	CustomerID int
	Itens      []*Itens
	CreatedAt  time.Time `gorm:"not null;autoCreateTime"`
	UpdatedAt  time.Time `gorm:"not null;autoUpdateTime"`
}

type OrderItens struct {
	ID        int       `gorm:"not null"`
	OrderID   int       `gorm:"not null"`
	Itens     []*Itens  `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null;autoCreateTime"`
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime"`
}

type Itens struct {
	ProductID int `gorm:"not null"`
	Quantity  int `gorm:"not_null"`
}

// OrderResponseList summary list
type OrderResponseList struct {
	Result []*Order
	Count  int64
}
