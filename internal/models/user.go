package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name           string  `json:"name"`
	Username       string  `json:"username" gorm:"unique"`
	Email          string  `json:"email" gorm:"unique"`
	PasswordHash   string  `json:"-"`
	PhoneNumber    string  `json:"phone_number" gorm:"unique"`
	ProfilePicture string  `json:"profile_picture"`
	Bio            string  `json:"bio"`
	Language       string  `json:"language"`
	IsVerified     bool    `json:"is_verified"`
	Events         []Event `gorm:"foreignKey:OrganizerID"`
}
