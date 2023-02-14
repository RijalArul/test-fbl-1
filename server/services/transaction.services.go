package services

import (
	"test-fbl-1/server/entities"
	"test-fbl-1/server/repositories"
	"test-fbl-1/server/webs"
)

type TransactionService interface {
	Create(transactionDTO webs.TransactionDTO, userID uint) (interface{}, error)
	FindAll() ([]webs.TransactionResponse, error)
}

type TransactionServiceImpl struct {
	transactionRepository repositories.TransactionRepository
	companyRepository     repositories.CompanyRepository
	productRepository     repositories.ProductRepository
}

func NewTransactionService(TransactionRepository repositories.TransactionRepository, CompanyRepository repositories.CompanyRepository, ProductRepository repositories.ProductRepository) TransactionService {
	return &TransactionServiceImpl{
		transactionRepository: TransactionRepository,
		companyRepository:     CompanyRepository,
		productRepository:     ProductRepository,
	}
}

func TransactionResponseBody(transaction *entities.Transaction) webs.TransactionResponse {
	return webs.TransactionResponse{
		CreatedAt:     transaction.CreatedAt,
		CompanyName:   transaction.CompanyName,
		ProductName:   transaction.ProductName,
		TotalQuantity: transaction.TotalQuantity,
		Price:         transaction.Price,
		TotalPrice:    transaction.TotalPrice,
		RestStock:     transaction.Product.Stock - transaction.TotalQuantity,
	}
}

func (s *TransactionServiceImpl) Create(transactionDTO webs.TransactionDTO, userID uint) (interface{}, error) {
	company, err := s.companyRepository.FindByID(transactionDTO.CompanyID)
	product, err := s.productRepository.FindByID(transactionDTO.ProductID)

	transaction := entities.Transaction{
		CompanyName:   company.Name,
		ProductName:   product.Name,
		TotalQuantity: transactionDTO.TotalQuantity,
		Price:         product.Price,
		TotalPrice:    product.Price * transactionDTO.TotalQuantity,
		UserID:        userID,
		CompanyID:     company.ID,
		ProductID:     product.ID,
	}

	createTransaction, err := s.transactionRepository.Create(transaction, *product, product.ID)
	transactionBody := TransactionResponseBody(createTransaction)
	return transactionBody, err
}

func (s *TransactionServiceImpl) FindAll() ([]webs.TransactionResponse, error) {
	transactions, err := s.transactionRepository.FindAll()

	newTransactions := []webs.TransactionResponse{}

	for i := 0; i < len(transactions); i++ {
		transaction := TransactionResponseBody(&transactions[i])
		newTransactions = append(newTransactions, transaction)
	}

	return newTransactions, err
}
