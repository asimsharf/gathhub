package models

import "gorm.io/gorm"

type UserBadge struct {
	gorm.Model
	UserID    uint   `json:"user_id"`
	BadgeType string `json:"badge_type"`
	EarnedAt  string `json:"earned_at"`
}
