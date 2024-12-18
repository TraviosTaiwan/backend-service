package middlewares

import (
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/iwachan14736/travios-backend-service/pkg/config"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		if token == "" {
			return c.JSON(http.StatusUnauthorized, "Missing token")
		}

		claims := jwt.MapClaims{}
		_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.LocalConfig.JWTSecret), nil
		})

		if err != nil {
			return c.JSON(http.StatusUnauthorized, "Invalid token")
		}

		c.Set("user", claims)
		return next(c)
	}
}
