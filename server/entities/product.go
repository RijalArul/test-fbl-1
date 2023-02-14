package entities

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Product struct {
	GormModel
	Name      string `gorm:"not null;unique" valid:"required~Your Products Name is required"`
	Slug      string `gorm:"not null" valid:"required~Your Products Slug Name is required"`
	Price     int    `gorm:"not null" valid:"required~Your Price Products is required"`
	Stock     int    `gorm:"not null" valid:"required~Your Product Stock is required"`
	UserID    uint
	CompanyID uint
	User      *User
	Company   *Company
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
