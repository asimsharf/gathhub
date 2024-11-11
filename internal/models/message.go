package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	GroupID        uint   `json:"group_id"`
	EventID        uint   `json:"event_id"`
	SenderID       uint   `json:"sender_id"`
	MessageContent string `json:"message_content"`
	SentAt         string `json:"sent_at"`
}
