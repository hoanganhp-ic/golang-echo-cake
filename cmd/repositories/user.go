package repositories

import "fitness-api/cmd/models"

type UserRepositories interface {
	GetByEmail(email string) (*models.User, error)
	Create(user *models.User) error
	GetByID(id uint) (*models.User, error)
}
