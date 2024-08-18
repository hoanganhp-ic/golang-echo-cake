package repositoryimpl

import (
	"fitness-api/cmd/dto"
	"fitness-api/cmd/models"
	"fitness-api/cmd/repositories"
	"fitness-api/cmd/utils"
	"strings"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type CakeRepositoryImpl struct {
	db *gorm.DB
}

var _ repositories.CakeRepositories = (*CakeRepositoryImpl)(nil)

func NewCakeRepositoryImpl(db *gorm.DB) repositories.CakeRepositories {
	return &CakeRepositoryImpl{db: db}
}

func (cr *CakeRepositoryImpl) Create(cake models.Cake) error {
	err := cr.db.Create(&cake).Error
	if err != nil {
		log.Errorf("Cannot create cake: %s", err)
		return err
	}
	return nil
}

func (cr *CakeRepositoryImpl) Search(searchCake dto.SearchCake) ([]models.Cake, error) {
	var cakes []models.Cake
	paginate := utils.NewPaginate(searchCake.PageSize, searchCake.Page)
	err := cr.db.Debug().Where("LOWER(name) LIKE ?", "%"+strings.ToLower(searchCake.Name)+"%").Where("user_id = ?", searchCake.UserID).Scopes(paginate.PaginatedResult).Find(&cakes).Error
	if err != nil {
		log.Errorf("not found: %s", err)
		return cakes, err
	}
	return cakes, nil
}

func (cr *CakeRepositoryImpl) GetByID(id int) (models.Cake, error) {
	var cake models.Cake
	err := cr.db.First(&cake, id).Error
	if err != nil {
		log.Errorf("cake %s", err)
		return cake, err
	}
	return cake, nil
}

func (cr *CakeRepositoryImpl) DeleteByID(id int) error {
	err := cr.db.Delete(&models.Cake{}, id).Error
	if err != nil {
		log.Errorf("cannot Delete cakeId: %d", id)
		return err
	}
	return nil
}

func (cr *CakeRepositoryImpl) UpdateByID(id int, cake models.Cake) error {
	err := cr.db.Model(&models.Cake{}).Where("id = ?", id).Where("user_id = ?", cake.UserID).Updates(cake).Error
	if err != nil {
		log.Errorf("cannot update cakeId: %d", id)
		return err
	}
	return nil
}
