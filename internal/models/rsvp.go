package models

import "gorm.io/gorm"

type Rsvp struct {
	gorm.Model
	EventID  uint   `json:"event_id"`
	UserID   uint   `json:"user_id"`
	Status   string `json:"status"`
	RSVPDate string `json:"rsvp_date"`
}
