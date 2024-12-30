package repositories

import (
	"github.com/iwachan14736/travios-backend-service/pkg/domain"
	"github.com/iwachan14736/travios-backend-service/pkg/models"
	"gorm.io/gorm"
)

type itemRepo struct {
	db *gorm.DB
}

func ItemDBInstance(d *gorm.DB) domain.IItemRepository {
	return &itemRepo{
		db: d,
	}
}

func (repo *itemRepo) CreateItem(item *models.Item) error {
	if err := repo.db.Create(item).Error; err != nil {
		return err
	}
	return nil
}

func (repo *itemRepo) GetItems(itemID uint) []models.Item {
	var items []models.Item
	var err error

	if itemID != 0 {
		err = repo.db.Where("id = ?", itemID).Find(&items).Error
	} else {
		err = repo.db.Find(&items).Error
	}

	if err != nil {
		return []models.Item{}
	}
	return items
}

func (repo *itemRepo) GetItemByTag(tag string) ([]models.Item, error) {
	var items []models.Item
	if err := repo.db.Where("tag LIKE ?", "%"+tag+"%").Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (repo *itemRepo) UpdateItem(item *models.Item) error {
	if err := repo.db.Save(item).Error; err != nil {
		return err
	}
	return nil
}

func (repo *itemRepo) DeleteItem(itemID uint) error {
	var item models.Item
	if err := repo.db.
		Where("id = ?", itemID).Delete(&item).
		Error; err != nil {
		return err
	}
	return nil
}
