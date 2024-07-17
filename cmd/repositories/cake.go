package repositories

import (
	"fitness-api/cmd/models"
	"fitness-api/cmd/storage"
	"log"
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
