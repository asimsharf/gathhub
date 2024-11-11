package services

import (
    "github.com/asimsharf/gatherhub/internal/models"
    "github.com/asimsharf/gatherhub/internal/repositories"
)

// NotificationService defines the business logic for Notification
type NotificationService struct {
    Repo *repositories.NotificationRepository
}

func NewNotificationService(repo *repositories.NotificationRepository) *NotificationService {
    return &NotificationService{Repo: repo}
}

func (s *NotificationService) CreateNotification(item *models.Notification) error {
    return s.Repo.CreateNotification(item)
}

func (s *NotificationService) GetNotificationByID(id uint) (*models.Notification, error) {
    return s.Repo.GetNotificationByID(id)
}

func (s *NotificationService) GetAllNotifications() ([]models.Notification, error) {
    return s.Repo.GetAllNotifications()
}

func (s *NotificationService) UpdateNotification(item *models.Notification) error {
    return s.Repo.UpdateNotification(item)
}

func (s *NotificationService) DeleteNotification(id uint) error {
    return s.Repo.DeleteNotification(id)
}
