package repository

import (
	"context"
	"soat1-challenge1/internal/core/domain"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Product Category struct type
type ProductCategory struct {
	db *gorm.DB
}

// NewProductCategoryRepository instantiates product repository
func NewProductCategoryRepository(db *gorm.DB) *ProductCategory {
	return &ProductCategory{db}
}

// CreateCategory create a new category record. If record already exists, it does nothing.
func (p *ProductCategory) CreateCategory(ctx context.Context, c *domain.Category) (*domain.Category, error) {
	dbFn := p.db.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true})

	if result := dbFn.Table("product_category").Create(c); result.Error != nil {
		return nil, result.Error
	}

	return c, nil
}

// List retrives all categories.
func (o *ProductCategory) GetCategories(ctx context.Context) (*domain.CategoriesResponseList, error) {
	dbFn := o.db.WithContext(ctx)

	var count int64
	var categories []*domain.Category

	result := dbFn.Table("product_category").Find(&categories).Count(&count)

	if result.Error != nil {
		return nil, result.Error
	}
	return &domain.CategoriesResponseList{
		Result: categories,
		Count:  count,
	}, nil
}

// List retrives all category products.
func (o *ProductCategory) GetCategoryProducts(ctx context.Context, category *domain.Category) (*domain.ProductResponseList, error) {
	dbFn := o.db.WithContext(ctx)
	categoryID := category.ID

	var count int64
	var products []*domain.Product

	result := dbFn.Table("product").Where("category_id = ?", categoryID).Find(&products).Count(&count)

	if result.Error != nil {
		return nil, result.Error
	}
	return &domain.ProductResponseList{
		Result: products,
		Count:  count,
	}, nil
}
