package services

import (
	"test-fbl-1/server/entities"
	"test-fbl-1/server/repositories"
	"test-fbl-1/server/webs"

	"github.com/gosimple/slug"
)

type ProductService interface {
	Create(productDTO webs.ProductDTO, userID uint) (webs.ProductBodyResponse, error)
}

type ProductServiceImpl struct {
	productRepository repositories.ProductRepository
	companyRepository repositories.CompanyRepository
}

func NewProductService(ProductRepository repositories.ProductRepository, CompanyRepository repositories.CompanyRepository) ProductService {
	return &ProductServiceImpl{productRepository: ProductRepository, companyRepository: CompanyRepository}
}

func ProductCreateBodyResponse(product *entities.Product) webs.ProductBodyResponse {
	return webs.ProductBodyResponse{
		ProductName:  product.Name,
		ProductSlug:  product.Slug,
		ProductPrice: product.Price,
		ProductStock: product.Stock,
		Company:      product.Company,
	}
}

func (s *ProductServiceImpl) Create(productDTO webs.ProductDTO, userID uint) (webs.ProductBodyResponse, error) {
	company, err := s.companyRepository.FindByID(productDTO.CompanyID)
	product := entities.Product{
		Name:      productDTO.ProductName,
		Price:     productDTO.ProductPrice,
		Slug:      slug.Make(productDTO.ProductName),
		Stock:     productDTO.ProductStock,
		UserID:    userID,
		CompanyID: company.ID,
	}

	createCompany, err := s.productRepository.Create(product)
	respCompany := ProductCreateBodyResponse(createCompany)

	return respCompany, err
}
