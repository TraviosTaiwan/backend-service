package controllers

import (
	"net/http"

	"github.com/iwachan14736/travios-backend-service/pkg/domain"
	"github.com/iwachan14736/travios-backend-service/pkg/types"
	"github.com/labstack/echo/v4"
)

var AuthService domain.IAuthService

func SetAuthService(aService domain.IAuthService) {
	AuthService = aService
}

// @Summary Login user
// @Description Authenticate a user and return JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body types.LoginRequest true "Login Credentials"
// @Success 200 {object} types.AuthResponse
// @Failure 400 {string} string "Invalid request"
// @Failure 401 {string} string "Invalid credentials"
// @Router /login [post]
func Login(c echo.Context) error {
	req := new(types.LoginRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	resp, err := AuthService.Login(req)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}

// @Summary Register new user
// @Description Register a new user and return JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param user body types.RegisterRequest true "User Registration Details"
// @Success 201 {object} types.AuthResponse
// @Failure 400 {string} string "Invalid request"
// @Router /register [post]
func Register(c echo.Context) error {
	req := new(types.RegisterRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	resp, err := AuthService.Register(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, resp)
}

// @Summary Get anonymous token
// @Description Generate a temporary anonymous token valid for 1 hour
// @Tags auth
// @Produce json
// @Success 200 {object} map[string]string
// @Router /anonymous-token [get]
func GetAnonymousToken(c echo.Context) error {
	token, err := AuthService.GenerateAnonymousToken()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to generate token")
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}
