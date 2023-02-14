package entities

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Transaction struct {
	GormModel
	CompanyName   string `gorm:"not null;" valid:"required~Your Transactions CompanyName is required"`
	ProductName   string `gorm:"not null" valid:"required~Your Transactions ProductName is required"`
	TotalQuantity int    `gorm:"not null" valid:"required~Your TotalQuantity Transactions is required"`
	Price         int    `gorm:"not null" valid:"required~Your Transaction Price is required"`
	TotalPrice    int    `gorm:"not null" valid:"required~Your Transaction TotalPrice is required"`
	UserID        uint
	CompanyID     uint
	ProductID     uint
	User          *User
	Company       *Company
	Product       *Product
}

func (t *Transaction) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(t)

	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}
