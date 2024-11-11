package repositories

import (
    "gorm.io/gorm"
    "github.com/asimsharf/gatherhub/internal/models"
)

// GroupRepository defines methods for CRUD operations on Group
type GroupRepository struct {
    db *gorm.DB
}

func NewGroupRepository(db *gorm.DB) *GroupRepository {
    return &GroupRepository{db: db}
}

func (r *GroupRepository) CreateGroup(item *models.Group) error {
    return r.db.Create(item).Error
}

func (r *GroupRepository) GetGroupByID(id uint) (*models.Group, error) {
    var item models.Group
    if err := r.db.First(&item, "id = ?", id).Error; err != nil {
        return nil, err
    }
    return &item, nil
}

func (r *GroupRepository) GetAllGroups() ([]models.Group, error) {
    var items []models.Group
    if err := r.db.Find(&items).Error; err != nil {
        return nil, err
    }
    return items, nil
}

func (r *GroupRepository) UpdateGroup(item *models.Group) error {
    return r.db.Save(item).Error
}

func (r *GroupRepository) DeleteGroup(id uint) error {
    return r.db.Delete(&models.Group{}, id).Error
}
