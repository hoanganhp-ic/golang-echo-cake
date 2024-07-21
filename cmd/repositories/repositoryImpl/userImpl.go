package repositoryimpl

import (
	"fitness-api/cmd/models"
	"fitness-api/cmd/repositories"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

var _ repositories.UserRepositories = (*UserRepositoryImpl)(nil)

func NewUserRepositoryImpl(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

func (ur *UserRepositoryImpl) Create(user *models.User) error {
	return ur.db.Create(user).Error
}

func (ur *UserRepositoryImpl) GetByEmail(email string) (*models.User, error) {
	var user models.User
	if err := ur.db.Where(&models.User{Email: email}).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
