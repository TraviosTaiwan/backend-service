package services

import (
	"errors"

	"github.com/iwachan14736/travios-backend-service/pkg/domain"
	"github.com/iwachan14736/travios-backend-service/pkg/models"
	"github.com/iwachan14736/travios-backend-service/pkg/types"
)

type CategoryService struct {
	repo domain.ICategoryRepository
}

func CategoryServiceInstance(categoryRepo domain.ICategoryRepository) domain.ICategoryService {
	return &CategoryService{
		repo: categoryRepo,
	}
}

func (service *CategoryService) GetCategories(categoryID uint) ([]types.Category, error) {
	var allCategories []types.Category
	categories := service.repo.GetCategories(categoryID)
	if len(categories) == 0 {
		return nil, errors.New("Category not found")
	}
	for _, val := range categories {
		allCategories = append(allCategories, types.Category{
			ID:          val.ID,
			Name:        val.Name,
			Description: val.Description,
		})
	}
	return allCategories, nil
}

func (service *CategoryService) CreateCategory(categoryRequest *models.Category) error {
	return service.repo.CreateCategory(categoryRequest)
}

func (service *CategoryService) UpdateCategory(categoryRequest *models.Category) error {
	return service.repo.UpdateCategory(categoryRequest)
}

func (service *CategoryService) DeleteCategory(categoryID uint) error {
	return service.repo.DeleteCategory(categoryID)
}
