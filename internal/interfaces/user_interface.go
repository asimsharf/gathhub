package interfaces

import (
	"github.com/asimsharf/gatherhub/internal/models"
)

// UserService defines the interface for User service methods
type UserService interface {
	CreateUser(item *models.User) error
	GetUserByID(id uint) (*models.User, error)
	GetAllUsers() ([]models.User, error)
	UpdateUser(item *models.User) error
	DeleteUser(id uint) error
}
