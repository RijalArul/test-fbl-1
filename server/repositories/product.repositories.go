package repositories

import (
	"test-fbl-1/server/entities"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductRepository interface {
	Create(product entities.Product) (*entities.Product, error)
	FindAll(companyID uint) ([]entities.Product, error)
	FindByID(productID uint) (*entities.Product, error)
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

func (r *ProductRepositoryImpl) FindAll(companyID uint) ([]entities.Product, error) {
	var products []entities.Product
	err := r.db.Model(products).Where("company_id = ?", companyID).Find(products).Error
	return products, err
}

func (r *ProductRepositoryImpl) FindByID(ProductID uint) (*entities.Product, error) {
	var product entities.Product
	err := r.db.Preload(clause.Associations).Model(product).Where("id = ?", ProductID).First(&product).Error
	return &product, err
}
