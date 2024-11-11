package interfaces

import (
    "github.com/asimsharf/gatherhub/internal/models"
)

// RsvpService defines the interface for Rsvp service methods
type RsvpService interface {
    CreateRsvp(item *models.Rsvp) error
    GetRsvpByID(id uint) (*models.Rsvp, error)
    GetAllRsvps() ([]models.Rsvp, error)
    UpdateRsvp(item *models.Rsvp) error
    DeleteRsvp(id uint) error
}
