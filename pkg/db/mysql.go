package db

import (
	"log"

	"github.com/asimsharf/gatherhub/config"
	"github.com/asimsharf/gatherhub/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectMySQL() (*gorm.DB, string) {
	// Load the configuration
	cfg := config.LoadConfig()

	// Establish a database connection
	db, err := gorm.Open(mysql.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	// Perform auto-migration
	err = db.AutoMigrate(
		&models.User{},
		&models.Event{},
		&models.Group{},
		&models.GroupMember{},
		&models.Rsvp{},
		&models.Message{},
		&models.BusinessPartner{},
		&models.SponsoredEvent{},
		&models.UserBadge{},
		&models.Notification{},
		&models.EventCategory{},
		&models.EventCategoryMap{},
		&models.Language{},
		&models.Translation{},
	)
	if err != nil {
		log.Fatalf("Could not migrate the database: %v", err)
	}

	return db, cfg.AppPort
}
