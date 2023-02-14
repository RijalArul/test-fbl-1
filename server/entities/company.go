package entities

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Company struct {
	GormModel
	Name     string `gorm:"not null;unique" valid:"required~Your Company Name is required"`
	Code     string `gorm:"not null;unique" valid:"required~Your Company Code is required"`
	UserID   uint
	Products []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	User     *User
	Product  *Product
}

func (c *Company) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(c)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
