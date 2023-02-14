package repositories

import (
	"test-fbl-1/server/entities"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TransactionRepository interface {
	Create(transaction entities.Transaction, product entities.Product, productID uint) (*entities.Transaction, error)
	FindAll() ([]entities.Transaction, error)
}

type TransactionRepositoryImpl struct {
	db *gorm.DB
}

func NewTransactionRepository(DB *gorm.DB) TransactionRepository {
	return &TransactionRepositoryImpl{db: DB}
}

func (r *TransactionRepositoryImpl) Create(transaction entities.Transaction, product entities.Product, productID uint) (*entities.Transaction, error) {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {

		return nil, err
	}

	transaction.ProductID = productID

	if err := tx.Preload(clause.Associations).Create(&transaction).First(&transaction).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	product.Stock = product.Stock - transaction.TotalQuantity

	if err := tx.Preload(clause.Associations).Where("id = ? ", productID).Updates(&product).First(&product).Error; err != nil {
		tx.Rollback()
	}

	return &transaction, tx.Commit().Error
}

func (r *TransactionRepositoryImpl) FindAll() ([]entities.Transaction, error) {
	var transactions []entities.Transaction
	err := r.db.Preload(clause.Associations).Model(transactions).First(&transactions).Error

	return transactions, err
}
