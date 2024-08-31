package repositories

import "fitness-api/cmd/models"

type CategoryRepositories interface {
	GetAll() ([]models.Category, error)
	GetByID(id int) (models.Category, error)
	Create(category *models.Category) (models.Category, error)
	Update(id int, category models.Category) error
	GetByName(name string) (models.Category, error)
	ExistsByName(name string) (bool, error)
}
