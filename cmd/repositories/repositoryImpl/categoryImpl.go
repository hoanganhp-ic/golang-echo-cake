package repositoryimpl

import (
	"fitness-api/cmd/models"
	"fitness-api/cmd/repositories"

	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
	db *gorm.DB
}

var _ repositories.CategoryRepositories = (*CategoryRepositoryImpl)(nil)

func NewCategoryRepositoryImpl(db *gorm.DB) repositories.CategoryRepositories {
	return &CategoryRepositoryImpl{db: db}
}

func (c *CategoryRepositoryImpl) GetAll() ([]models.Category, error) {
	var categories []models.Category
	if err := c.db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (c *CategoryRepositoryImpl) GetByID(id int) (models.Category, error) {
	var category models.Category
	if err := c.db.First(&category, id).Error; err != nil {
		return category, err
	}
	return category, nil
}

func (c *CategoryRepositoryImpl) Create(category *models.Category) (models.Category, error) {
	if err := c.db.Create(&category).Error; err != nil {
		return models.Category{}, err
	}
	return *category, nil
}

func (c *CategoryRepositoryImpl) Update(id int, category models.Category) error {
	if err := c.db.Model(&category).Where("id = ?", id).Updates(&category).Error; err != nil {
		return err
	}
	return nil
}

func (c *CategoryRepositoryImpl) GetByName(name string) (models.Category, error) {
	var category models.Category
	if err := c.db.Where("name = ?", name).First(&category).Error; err != nil {
		return category, err
	}
	return category, nil
}

func (c *CategoryRepositoryImpl) ExistsByName(name string) (bool, error) {
	var category models.Category
	if err := c.db.Where("name = ?", name).First(&category).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
