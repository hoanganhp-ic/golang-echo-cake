package models

import (
	"time"
)

type Cake struct {
	ID          int       `json:"id" gorm:"index;not null"`
	Name        string    `json:"name" gorm:"not null"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time // Automatically managed by GORM for creation time
	UpdatedAt   time.Time // Automatically managed by GORM for update time
}
