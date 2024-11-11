package models

import "gorm.io/gorm"

type EventCategory struct {
	gorm.Model
	CategoryName string `json:"category_name" gorm:"unique"`
	Description  string `json:"description"`
}
