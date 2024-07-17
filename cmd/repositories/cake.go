package repositories

import (
	"fitness-api/cmd/dto"
	"fitness-api/cmd/models"
	"fitness-api/cmd/storage"
	"fitness-api/cmd/until"
	"log"
	"strings"
)

func Create(cake models.Cake) (models.Cake, error) {
	db := storage.GetDB()
	err := db.Create(&cake).Error
	if err != nil {
		log.Fatal(err)
		return cake, err
	}
	return cake, nil
}

func GetAll() ([]models.Cake, error) {
	db := storage.GetDB()
	var cakes []models.Cake
	err := db.Find(&cakes).Error
	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
		return cakes, err
	}
	return cakes, nil
}
