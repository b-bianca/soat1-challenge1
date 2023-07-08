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

	if result := dbFn.Table("product_category").Create(c); result.Error != nil {
		return nil, result.Error
	}

	return c, nil
}

// Create create a new product record. If record already exists, it does nothing.
func (p *Product) Create(ctx context.Context, pdt *domain.Product) (*domain.Product, error) {
	dbFn := p.db.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true})

	if result := dbFn.Table("product").Create(pdt); result.Error != nil {
		return nil, result.Error
	}

	return pdt, nil
}

// Update updates a product record in database.
func (p *Product) Update(ctx context.Context, pdt *domain.Product) error {
	dbFn := p.db.WithContext(ctx)

	filter := pdt.ID

	return dbFn.Table("product").Where(filter).Updates(pdt).Error
}

// Delete remove a product record in database.
func (p *Product) Delete(ctx context.Context, pdt *domain.Product) error {
	dbFn := p.db.WithContext(ctx)

	id := pdt.ID

	return dbFn.Table("product").Where("id = ?", id).Delete(&pdt).Error
}
