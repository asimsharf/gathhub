package services

import (
    "github.com/asimsharf/gatherhub/internal/models"
    "github.com/asimsharf/gatherhub/internal/repositories"
)

// RsvpService defines the business logic for Rsvp
type RsvpService struct {
    Repo *repositories.RsvpRepository
}

func NewRsvpService(repo *repositories.RsvpRepository) *RsvpService {
    return &RsvpService{Repo: repo}
}

func (s *RsvpService) CreateRsvp(item *models.Rsvp) error {
    return s.Repo.CreateRsvp(item)
}

func (s *RsvpService) GetRsvpByID(id uint) (*models.Rsvp, error) {
    return s.Repo.GetRsvpByID(id)
}

func (s *RsvpService) GetAllRsvps() ([]models.Rsvp, error) {
    return s.Repo.GetAllRsvps()
}

func (s *RsvpService) UpdateRsvp(item *models.Rsvp) error {
    return s.Repo.UpdateRsvp(item)
}

func (s *RsvpService) DeleteRsvp(id uint) error {
    return s.Repo.DeleteRsvp(id)
}
