package interfaces

import (
	"github.com/asimsharf/gatherhub/internal/models"
)

// EventService defines the interface for Event service methods
type EventService interface {
	CreateEvent(item *models.Event) error
	GetEventByID(id uint) (*models.Event, error)
	GetAllEvents() ([]models.Event, error)
	UpdateEvent(item *models.Event) error
	DeleteEvent(id uint) error
}
