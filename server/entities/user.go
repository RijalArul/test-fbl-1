package entities

import (
	"test-fbl-1/server/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username string    `gorm:"not null;unique" valid:"required~Your username is required"`
	Password string    `gorm:"not null" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Role     string    `gorm:"not null" valid:"required~Your role is required"`
	Products []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)

	err = nil
	return
}
