package repositories

import (
	"github.com/iwachan14736/travios-backend-service/pkg/domain"
	"github.com/iwachan14736/travios-backend-service/pkg/models"
	"gorm.io/gorm"
)

type categoryRepo struct {
	db *gorm.DB
}

func CategoryDBInstance(d *gorm.DB) domain.ICategoryRepository {
	return &categoryRepo{
		db: d,
	}
}

func (repo *categoryRepo) CreateCategory(category *models.Category) error {
	if err := repo.db.Create(category).Error; err != nil {
		return err
	}
	return nil
}

func (repo *categoryRepo) GetCategories(categoryID uint) []models.Category {
	var categories []models.Category
	var err error

	if categoryID != 0 {
		err = repo.db.Where("id = ?", categoryID).Find(&categories).Error
	} else {
		err = repo.db.Find(&categories).Error
	}

	if err != nil {
		return []models.Category{}
	}
	return categories
}

func (repo *categoryRepo) UpdateCategory(category *models.Category) error {
	if err := repo.db.Save(category).Error; err != nil {
		return err
	}
	return nil
}

func (repo *categoryRepo) DeleteCategory(categoryID uint) error {
	var category models.Category
	if err := repo.db.
		Where("id = ?", categoryID).Delete(&category).
		Error; err != nil {
		return err
	}
	return nil
}
