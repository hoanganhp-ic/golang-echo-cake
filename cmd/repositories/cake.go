package repositories

import (
	"fitness-api/cmd/dto"
	"fitness-api/cmd/models"
)

type CakeRepositories interface {
	Create(cake models.Cake) error
	Search(searchCake dto.SearchCake) ([]models.Cake, error)
	GetByID(id int) (models.Cake, error)
	DeleteByID(id int) error
	UpdateByID(id int, cake models.Cake) error
}

// func Create(cake models.Cake) (models.Cake, error) {
// 	db := storage.GetDB()
// 	err := db.Create(&cake).Error
// 	if err != nil {
// 		log.Errorf("Cannot create cake: %s", err)
// 		return cake, err
// 	}
// 	return cake, nil
// }

// func GetAll() ([]models.Cake, error) {
// 	db := storage.GetDB()
// 	var cakes []models.Cake
// 	err := db.Find(&cakes).Error
// 	if err != nil {
// 		log.Errorf("Cannot get all cakes")
// 		return cakes, err
// 	}
// 	return cakes, nil
// }

// func Search(searchCake dto.SearchCake) ([]models.Cake, error) {
// 	db := storage.GetDB()
// 	var cakes []models.Cake
// 	paginate := utils.NewPaginate(searchCake.PageSize, searchCake.Page)
// 	err := db.Debug().Where("LOWER(name) LIKE ?", "%"+strings.ToLower(searchCake.Name)+"%").Scopes(paginate.PaginatedResult).Find(&cakes).Error
// 	if err != nil {
// 		log.Errorf("not found: %s", err)
// 		return cakes, err
// 	}
// 	return cakes, nil
// }

// func GetByID(id int) (models.Cake, error) {
// 	db := storage.GetDB()
// 	var cake models.Cake
// 	err := db.First(&cake, id).Error
// 	if err != nil {
// 		log.Errorf("cake %s", err)
// 		return cake, err
// 	}
// 	return cake, nil
// }

// func DeleteByID(id int) error {
// 	db := storage.GetDB()
// 	err := db.Delete(&models.Cake{}, id).Error
// 	if err != nil {
// 		log.Errorf("cannot Delete cakeId: %d", id)
// 		return err
// 	}
// 	return nil
// }
