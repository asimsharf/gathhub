package repositories

import (
    "gorm.io/gorm"
    "github.com/asimsharf/gatherhub/internal/models"
)

// NotificationRepository defines methods for CRUD operations on Notification
type NotificationRepository struct {
    db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) *NotificationRepository {
    return &NotificationRepository{db: db}
}

func (r *NotificationRepository) CreateNotification(item *models.Notification) error {
    return r.db.Create(item).Error
}

func (r *NotificationRepository) GetNotificationByID(id uint) (*models.Notification, error) {
    var item models.Notification
    if err := r.db.First(&item, "id = ?", id).Error; err != nil {
        return nil, err
    }
    return &item, nil
}

func (r *NotificationRepository) GetAllNotifications() ([]models.Notification, error) {
    var items []models.Notification
    if err := r.db.Find(&items).Error; err != nil {
        return nil, err
    }
    return items, nil
}

func (r *NotificationRepository) UpdateNotification(item *models.Notification) error {
    return r.db.Save(item).Error
}

func (r *NotificationRepository) DeleteNotification(id uint) error {
    return r.db.Delete(&models.Notification{}, id).Error
}
