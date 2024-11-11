package repositories

import (
    "gorm.io/gorm"
    "github.com/asimsharf/gatherhub/internal/models"
)

// RsvpRepository defines methods for CRUD operations on Rsvp
type RsvpRepository struct {
    db *gorm.DB
}

func NewRsvpRepository(db *gorm.DB) *RsvpRepository {
    return &RsvpRepository{db: db}
}

func (r *RsvpRepository) CreateRsvp(item *models.Rsvp) error {
    return r.db.Create(item).Error
}

func (r *RsvpRepository) GetRsvpByID(id uint) (*models.Rsvp, error) {
    var item models.Rsvp
    if err := r.db.First(&item, "id = ?", id).Error; err != nil {
        return nil, err
    }
    return &item, nil
}

func (r *RsvpRepository) GetAllRsvps() ([]models.Rsvp, error) {
    var items []models.Rsvp
    if err := r.db.Find(&items).Error; err != nil {
        return nil, err
    }
    return items, nil
}

func (r *RsvpRepository) UpdateRsvp(item *models.Rsvp) error {
    return r.db.Save(item).Error
}

func (r *RsvpRepository) DeleteRsvp(id uint) error {
    return r.db.Delete(&models.Rsvp{}, id).Error
}
