package routes

import (
	"github.com/iwachan14736/travios-backend-service/pkg/config"
	"github.com/iwachan14736/travios-backend-service/pkg/middlewares"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func InitRoutes(e *echo.Echo) {
	// Swagger route
	config.ConfigureSwagger()
	e.GET("/docs/*", echoSwagger.WrapHandler)

	traviosGroup := e.Group("/travios")

	// Auth routes (no middleware)
	HealthRoutes(traviosGroup)
	AuthRoutes(traviosGroup)

	// Protected routes
	protected := traviosGroup.Group("")
	protected.Use(middlewares.AuthMiddleware)

	CategoryRoutes(protected)
	VendorRoutes(protected)
	ItemRoutes(protected)
	CustomerRoutes(protected)
	SaleOrderRoutes(protected)
	SaleOrderItemRoutes(protected)
}
