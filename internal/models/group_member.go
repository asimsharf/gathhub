package models

type GroupMember struct {
	GroupID  uint   `gorm:"primaryKey"`
	UserID   uint   `gorm:"primaryKey"`
	JoinedAt string `json:"joined_at"`
}
