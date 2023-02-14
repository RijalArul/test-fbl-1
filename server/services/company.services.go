package services

import (
	"test-fbl-1/server/entities"
	"test-fbl-1/server/repositories"
	"test-fbl-1/server/webs"
)

type CompanyService interface {
	Create(companyDTO webs.CompanyDTO, userID uint) (webs.RepsonseCompanyBody, error)
	FindAll() ([]entities.Company, error)
}

type CompanyServiceImpl struct {
	companyRepository repositories.CompanyRepository
}

func NewCompanyService(CompanyRepository repositories.CompanyRepository) CompanyService {
	return &CompanyServiceImpl{companyRepository: CompanyRepository}
}

func RepsonseCompanyBody(company *entities.Company) webs.RepsonseCompanyBody {
	return webs.RepsonseCompanyBody{
		Name: company.Name,
		Code: company.Code,
	}
}

func (s *CompanyServiceImpl) Create(companyDTO webs.CompanyDTO, userID uint) (webs.RepsonseCompanyBody, error) {
	company := entities.Company{
		Name:   companyDTO.CompanyName,
		Code:   companyDTO.CompanyCode,
		UserID: userID,
	}

	createCompany, err := s.companyRepository.Create(company)
	respCompany := RepsonseCompanyBody(createCompany)

	return respCompany, err
}

func (s *CompanyServiceImpl) FindAll() ([]entities.Company, error) {
	companies, err := s.companyRepository.FindAll()
	return companies, err
}
