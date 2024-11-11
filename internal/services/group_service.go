package services

import (
    "github.com/asimsharf/gatherhub/internal/models"
    "github.com/asimsharf/gatherhub/internal/repositories"
)

// GroupService defines the business logic for Group
type GroupService struct {
    Repo *repositories.GroupRepository
}

func NewGroupService(repo *repositories.GroupRepository) *GroupService {
    return &GroupService{Repo: repo}
}

func (s *GroupService) CreateGroup(item *models.Group) error {
    return s.Repo.CreateGroup(item)
}

func (s *GroupService) GetGroupByID(id uint) (*models.Group, error) {
    return s.Repo.GetGroupByID(id)
}

func (s *GroupService) GetAllGroups() ([]models.Group, error) {
    return s.Repo.GetAllGroups()
}

func (s *GroupService) UpdateGroup(item *models.Group) error {
    return s.Repo.UpdateGroup(item)
}

func (s *GroupService) DeleteGroup(id uint) error {
    return s.Repo.DeleteGroup(id)
}
