package models

import "gorm.io/gorm"

type Notification struct {
	gorm.Model
	UserID  uint   `json:"user_id"`
	Message string `json:"message"`
	Type    string `json:"type"`
	IsRead  bool   `json:"is_read"`
}
