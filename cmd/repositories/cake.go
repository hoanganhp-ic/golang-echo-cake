package repositories

import (
	"fitness-api/cmd/dto"
	"fitness-api/cmd/models"
	"fitness-api/cmd/storage"
	"fitness-api/cmd/until"
	"strings"

	"github.com/labstack/gommon/log"
)

func Create(cake models.Cake) (models.Cake, error) {
	db := storage.GetDB()
	err := db.Create(&cake).Error
	if err != nil {
		log.Errorf("Cannot create cake: %s",err)
		return cake, err
	}
	return cake, nil
}

func GetAll() ([]models.Cake, error) {
	db := storage.GetDB()
	var cakes []models.Cake
	err := db.Find(&cakes).Error
	if err != nil {
		log.Errorf("Cannot get all cakes")
		return cakes, err
	}
	return cakes, nil
}

func Search(searchCake dto.SearchCake) ([]models.Cake, error) {
	db := storage.GetDB()
	var cakes []models.Cake
	paginate := until.NewPaginate(searchCake.PageSize, searchCake.Page)
	err := db.Debug().Where("LOWER(name) LIKE ?", "%"+strings.ToLower(searchCake.Name)+"%").Scopes(paginate.PaginatedResult).Find(&cakes).Error
	if err != nil {
		log.Errorf("not found: %s", err)
		return cakes, err
	}
	return cakes, nil
}

func GetByID(id int) (models.Cake, error) {
	db := storage.GetDB()
	var cake models.Cake
	err := db.First(&cake, id).Error
	if err != nil {
		log.Errorf("cake not found %s",err)
		return cake, err
	}
	return cake, nil
}

func DeleteByID(id int) error {
	db := storage.GetDB()
	err := db.Delete(&models.Cake{}, id).Error
	if err != nil {
		log.Errorf("cannot Delete cakeId: %d", id)
		return err
	}
	return nil
}
