package models

import "gorm.io/gorm"

type SponsoredEvent struct {
	gorm.Model
	EventID            uint   `json:"event_id"`
	PartnerID          uint   `json:"partner_id"`
	SponsorshipType    string `json:"sponsorship_type"`
	SponsorshipDetails string `json:"sponsorship_details"`
}
