package domain

import (
	"github.com/iwachan14736/travios-backend-service/pkg/models"
	"github.com/iwachan14736/travios-backend-service/pkg/types"
)

type ICategoryRepository interface {
	GetCategories(categoryID uint) []models.Category
	CreateCategory(category *models.Category) error
	UpdateCategory(category *models.Category) error
	DeleteCategory(categoryID uint) error
}

type ICategoryService interface {
	GetCategories(categoryID uint) ([]types.Category, error)
	CreateCategory(category *models.Category) error
	UpdateCategory(category *models.Category) error
	DeleteCategory(categoryID uint) error
}
