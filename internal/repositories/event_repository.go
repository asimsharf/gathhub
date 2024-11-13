package repositories

import (
	"github.com/asimsharf/gatherhub/internal/models"
	"gorm.io/gorm"
)

// EventRepository defines methods for CRUD operations on Event
type EventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) *EventRepository {
	return &EventRepository{db: db}
}

func (r *EventRepository) CreateEvent(item *models.Event) error {
	return r.db.Create(item).Error
}

func (r *EventRepository) GetEventByID(id uint) (*models.Event, error) {
	var item models.Event
	if err := r.db.First(&item, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *EventRepository) GetAllEvents() ([]models.Event, error) {
	var items []models.Event
	if err := r.db.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (r *EventRepository) UpdateEvent(item *models.Event) error {
	return r.db.Save(item).Error
}

func (r *EventRepository) DeleteEvent(id uint) error {
	return r.db.Delete(&models.Event{}, id).Error
}
