package repositories

import (
	"test-fbl-1/server/entities"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductRepository interface {
	Create(product entities.Product) (*entities.Product, error)
	FindAll() ([]entities.Product, error)
}

type ProductRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepository(DB *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{db: DB}
}

func (r *ProductRepositoryImpl) Create(product entities.Product) (*entities.Product, error) {
	err := r.db.Preload(clause.Associations).Create(&product).First(&product).Error
	return &product, err
}

func (r *ProductRepositoryImpl) FindAll() ([]entities.Product, error) {
	var products []entities.Product
	err := r.db.Model(products).Find(products).Error
	return products, err
}
