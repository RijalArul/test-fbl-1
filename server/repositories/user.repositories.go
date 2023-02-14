package repositories

import (
	databases "test-fbl-1/server/db"
	"test-fbl-1/server/entities"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user entities.User) (*entities.User, error)
	// FindByID()
	FindByUsername(username string) (*entities.User, error)
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(DB *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db: DB}
}

func (r *UserRepositoryImpl) Create(user entities.User) (*entities.User, error) {
	tx := databases.GetDB().Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {

		return nil, err
	}

	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &user, tx.Commit().Error
}

func (r *UserRepositoryImpl) FindByUsername(username string) (*entities.User, error) {
	var user entities.User
	err := r.db.Model(user).Where("username = ?", username).First(&user).Error

	return &user, err
}
