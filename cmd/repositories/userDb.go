package repositories

import (
	"fitness-api/cmd/models"
	"fitness-api/cmd/storage"
	"time"
)

func CreateUser(user models.User) (models.User, error) {
	db := storage.GetDB()
	sqlStatement := `INSERT INTO users (name, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := db.QueryRow(sqlStatement, user.Name, user.Email, user.Password, time.Now(), time.Now()).Scan(&user.Id)
	if err != nil {
		return user, err
	}
	return user, nil
}
