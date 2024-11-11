package services

import (
    "github.com/asimsharf/gatherhub/internal/models"
    "github.com/asimsharf/gatherhub/internal/repositories"
)

// EventService defines the business logic for Event
type EventService struct {
    Repo *repositories.EventRepository
}

func NewEventService(repo *repositories.EventRepository) *EventService {
    return &EventService{Repo: repo}
}

func (s *EventService) CreateEvent(item *models.Event) error {
    return s.Repo.CreateEvent(item)
}

func (s *EventService) GetEventByID(id uint) (*models.Event, error) {
    return s.Repo.GetEventByID(id)
}

func (s *EventService) GetAllEvents() ([]models.Event, error) {
    return s.Repo.GetAllEvents()
}

func (s *EventService) UpdateEvent(item *models.Event) error {
    return s.Repo.UpdateEvent(item)
}

func (s *EventService) DeleteEvent(id uint) error {
    return s.Repo.DeleteEvent(id)
}
