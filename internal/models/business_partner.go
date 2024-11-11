package models

import "gorm.io/gorm"

type BusinessPartner struct {
	gorm.Model
	BusinessName     string `json:"business_name"`
	ContactPerson    string `json:"contact_person"`
	Email            string `json:"email" gorm:"unique"`
	PhoneNumber      string `json:"phone_number"`
	Address          string `json:"address"`
	PartnershipStart string `json:"partnership_start"`
}
