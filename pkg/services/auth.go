package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/iwachan14736/travios-backend-service/pkg/config"
	"github.com/iwachan14736/travios-backend-service/pkg/constants"
	"github.com/iwachan14736/travios-backend-service/pkg/domain"
	"github.com/iwachan14736/travios-backend-service/pkg/models"
	"github.com/iwachan14736/travios-backend-service/pkg/types"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo domain.IUserRepository
}

func AuthServiceInstance(userRepo domain.IUserRepository) domain.IAuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

func (s *AuthService) Login(req *types.LoginRequest) (*types.AuthResponse, error) {
	user, err := s.userRepo.GetByUsername(req.Username)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	token, err := s.GenerateToken(user)
	if err != nil {
		return nil, err
	}

	return &types.AuthResponse{
		Token: token,
		User:  *user,
	}, nil
}

func (s *AuthService) Register(req *types.RegisterRequest) (*types.AuthResponse, error) {
	// Check if username exists
	if _, err := s.userRepo.GetByUsername(req.Username); err == nil {
		return nil, errors.New("username already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username: req.Username,
		Password: string(hashedPassword),
		Email:    req.Email,
		RoleID:   constants.RoleEmployee,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	token, err := s.GenerateToken(user)
	if err != nil {
		return nil, err
	}

	return &types.AuthResponse{
		Token: token,
		User:  *user,
	}, nil
}

func (s *AuthService) GenerateToken(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"role_id":  user.RoleID,
		"exp":      time.Now().Add(time.Hour * time.Duration(24)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.LocalConfig.JWTSecret))
}

func (s *AuthService) ValidateToken(tokenString string) (bool, error) {
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(config.LocalConfig.JWTSecret), nil
	})

	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *AuthService) GenerateAnonymousToken() (string, error) {
	claims := jwt.MapClaims{
		"anonymous": true,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.LocalConfig.JWTSecret))
}
