package models

import (
	"gorm.io/gorm"
)

type Cake struct {
	gorm.Model
	Name        string  `json:"name" gorm:"not null"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageUrl	   string  `json:"image_url"`
	UserID      uint     `json:"user_id"`
	User        User    `json:"user"`
	CategoryID  uint     `json:"category_id"`
}
