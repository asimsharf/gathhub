package interfaces

import (
    "github.com/asimsharf/gatherhub/internal/models"
)

// NotificationService defines the interface for Notification service methods
type NotificationService interface {
    CreateNotification(item *models.Notification) error
    GetNotificationByID(id uint) (*models.Notification, error)
    GetAllNotifications() ([]models.Notification, error)
    UpdateNotification(item *models.Notification) error
    DeleteNotification(id uint) error
}
