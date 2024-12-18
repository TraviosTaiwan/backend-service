package config

import (
	"github.com/iwachan14736/travios-backend-service/docs"
)

func ConfigureSwagger() {
	docs.SwaggerInfo.Title = "Travios Backend Service API"
	docs.SwaggerInfo.Description = "Travios Backend Service API documentation"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = LocalConfig.ApiHost
	docs.SwaggerInfo.BasePath = "/travios"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}
