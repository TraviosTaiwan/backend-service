package domain

import (
	"github.com/iwachan14736/travios-backend-service/pkg/models"
	"github.com/iwachan14736/travios-backend-service/pkg/types"
)

type IAuthService interface {
	Login(req *types.LoginRequest) (*types.AuthResponse, error)
	Register(req *types.RegisterRequest) (*types.AuthResponse, error)
	GenerateToken(user *models.User) (string, error)
	ValidateToken(token string) (bool, error)
}

type IUserRepository interface {
	GetByUsername(username string) (*models.User, error)
	Create(user *models.User) error
	Update(user *models.User) error
	Delete(id uint) error
	GetByID(id uint) (*models.User, error)
}
