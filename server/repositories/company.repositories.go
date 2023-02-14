package repositories

import (
	"test-fbl-1/server/entities"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CompanyRepository interface {
	Create(company entities.Company) (*entities.Company, error)
	FindAll() ([]entities.Company, error)
	FindByID(id uint) (*entities.Company, error)
}

type CompanyRepositoryImpl struct {
	db *gorm.DB
}

func NewCompanyRepository(DB *gorm.DB) CompanyRepository {
	return &CompanyRepositoryImpl{db: DB}
}

func (r *CompanyRepositoryImpl) Create(company entities.Company) (*entities.Company, error) {
	err := r.db.Create(&company).Error
	return &company, err
}

func (r *CompanyRepositoryImpl) FindAll() ([]entities.Company, error) {
	var companies []entities.Company
	err := r.db.Preload(clause.Associations).Model(companies).Find(&companies).Error
	return companies, err
}

func (r *CompanyRepositoryImpl) FindByID(id uint) (*entities.Company, error) {
	var company entities.Company
	err := r.db.Preload(clause.Associations).Model(company).Where("id = ?", id).First(&company).Error
	return &company, err
}
