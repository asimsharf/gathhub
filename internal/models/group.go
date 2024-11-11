package models

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Name         string `json:"name"`
	Description  string `json:"description"`
	CreatedBy    uint   `json:"created_by"`
	PrivacyLevel string `json:"privacy_level"`
	Members      []User `gorm:"many2many:group_members;"`
}
