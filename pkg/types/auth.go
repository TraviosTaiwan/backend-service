package types

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/iwachan14736/travios-backend-service/pkg/models"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r LoginRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Username, validation.Required, validation.Length(3, 50)),
		validation.Field(&r.Password, validation.Required, validation.Length(6, 100)),
	)
}

type RegisterRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (r RegisterRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Username, validation.Required, validation.Length(3, 50)),
		validation.Field(&r.Password, validation.Required, validation.Length(6, 100)),
		validation.Field(&r.Email, validation.Required, is.Email),
		validation.Field(&r.FirstName, validation.Required, validation.Length(2, 50)),
		validation.Field(&r.LastName, validation.Required, validation.Length(2, 50)),
	)
}

type AuthResponse struct {
	Token string      `json:"token"`
	User  models.User `json:"user"`
}
