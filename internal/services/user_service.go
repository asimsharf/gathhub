package services

import (
    "github.com/asimsharf/gatherhub/internal/models"
    "github.com/asimsharf/gatherhub/internal/repositories"
)

// UserService defines the business logic for User
type UserService struct {
    Repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
    return &UserService{Repo: repo}
}

func (s *UserService) CreateUser(item *models.User) error {
    return s.Repo.CreateUser(item)
}

func (s *UserService) GetUserByID(id uint) (*models.User, error) {
    return s.Repo.GetUserByID(id)
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
    return s.Repo.GetAllUsers()
}

func (s *UserService) UpdateUser(item *models.User) error {
    return s.Repo.UpdateUser(item)
}

func (s *UserService) DeleteUser(id uint) error {
    return s.Repo.DeleteUser(id)
}
