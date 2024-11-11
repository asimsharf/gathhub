package interfaces

import (
    "github.com/asimsharf/gatherhub/internal/models"
)

// GroupService defines the interface for Group service methods
type GroupService interface {
    CreateGroup(item *models.Group) error
    GetGroupByID(id uint) (*models.Group, error)
    GetAllGroups() ([]models.Group, error)
    UpdateGroup(item *models.Group) error
    DeleteGroup(id uint) error
}
