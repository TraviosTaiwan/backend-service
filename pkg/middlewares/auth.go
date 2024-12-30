package middlewares

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/iwachan14736/travios-backend-service/pkg/config"
	"github.com/labstack/echo/v4"
)

var anonymousEndpoints = map[string]bool{
	"/travios/sale_order/frontend": true,
	"/travios/item":                true,
}

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Check if endpoint allows anonymous access
		if anonymousEndpoints[c.Path()] {
			token := c.Request().Header.Get("Authorization")
			if token != "" {
				// If token is provided, validate it's an anonymous token
				claims := jwt.MapClaims{}
				_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
					return []byte(config.LocalConfig.JWTSecret), nil
				})

				if err == nil {
					if isAnonymous, ok := claims["anonymous"].(bool); ok && isAnonymous {
						c.Set("anonymous", true)
						return next(c)
					}
				}
			}
			// Allow access even without token for these endpoints
			return next(c)
		}

		// Regular auth check for other endpoints
		token := c.Request().Header.Get("Authorization")
		if token == "" {
			return c.JSON(http.StatusUnauthorized, "Missing token")
		}

		// Remove 'Bearer ' prefix if present
		token = strings.TrimPrefix(token, "Bearer ")

		claims := jwt.MapClaims{}
		_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.LocalConfig.JWTSecret), nil
		})

		if err != nil {
			return c.JSON(http.StatusUnauthorized, "Invalid token")
		}

		if isAnonymous, ok := claims["anonymous"].(bool); ok && isAnonymous {
			c.Set("anonymous", true)
		} else {
			c.Set("user", claims)
		}

		return next(c)
	}
}
