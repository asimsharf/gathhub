package models

import "gorm.io/gorm"

type Event struct {
	gorm.Model
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Location    string  `json:"location"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	StartTime   string  `json:"start_time"`
	EndTime     string  `json:"end_time"`
	OrganizerID uint    `json:"organizer_id"`
	Organizer   User    `gorm:"foreignKey:OrganizerID"`
	EventType   string  `json:"event_type"`
}
