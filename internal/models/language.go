package models

import "gorm.io/gorm"

type Language struct {
	gorm.Model
	LanguageCode string `json:"language_code" gorm:"unique"`
	LanguageName string `json:"language_name"`
	IsActive     bool   `json:"is_active"`
}
