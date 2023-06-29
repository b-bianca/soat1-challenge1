package repository

import (
	"context"
	"soat1-challenge1/internal/core/domain"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Product struct type
type Product struct {
	db *gorm.DB
}

// NewProductRepository instantiates product repository
func NewProductRepository(db *gorm.DB) *Product {
	return &Product{db}
}

// CreateCategory create a new category record. If record already exists, it does nothing.
func (p *Product) CreateCategory(ctx context.Context, c *domain.Category) (*domain.Category, error) {
	dbFn := p.db.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true})

	if result := dbFn.Table("categories").Create(c); result.Error != nil {
		return nil, result.Error
	}

	return c, nil
}

// CreateCategory create a new category record. If record already exists, it does nothing.
func (p *Product) Create(ctx context.Context, input *domain.Product) (*domain.Product, error) {
	return nil, nil
}
