package repositories

import (
    "gorm.io/gorm"
    "github.com/asimsharf/gatherhub/internal/models"
)

// UserRepository defines methods for CRUD operations on User
type UserRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
    return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(item *models.User) error {
    return r.db.Create(item).Error
}

func (r *UserRepository) GetUserByID(id uint) (*models.User, error) {
    var item models.User
    if err := r.db.First(&item, "id = ?", id).Error; err != nil {
        return nil, err
    }
    return &item, nil
}

func (r *UserRepository) GetAllUsers() ([]models.User, error) {
    var items []models.User
    if err := r.db.Find(&items).Error; err != nil {
        return nil, err
    }
    return items, nil
}

func (r *UserRepository) UpdateUser(item *models.User) error {
    return r.db.Save(item).Error
}

func (r *UserRepository) DeleteUser(id uint) error {
    return r.db.Delete(&models.User{}, id).Error
}
