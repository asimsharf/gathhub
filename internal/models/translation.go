package models

import "gorm.io/gorm"

type Translation struct {
	gorm.Model
	LanguageCode string `json:"language_code"`
	ContentType  string `json:"content_type"`
	ContentID    uint   `json:"content_id"`
	ContentField string `json:"content_field"`
	Translation  string `json:"translation"`
}
