package repositories

import (
	"fitness-api/cmd/models"
	"fitness-api/cmd/storage"
	"time"
)

func CreateMeasurement(measurement models.Measurements) (models.Measurements, error) {
	db := storage.GetDB()
	sqlStatement := `INSERT INTO measurements (user_id, weight, height, body_fat, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := db.QueryRow(sqlStatement, measurement.UserId, measurement.Weight, measurement.Height, measurement.BodyFat, time.Now()).Scan(&measurement.Id)
	if err != nil {
		return measurement, err
	}

	return measurement, nil
}
