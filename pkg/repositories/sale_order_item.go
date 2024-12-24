package repositories

import (
	"github.com/iwachan14736/travios-backend-service/pkg/domain"
	"github.com/iwachan14736/travios-backend-service/pkg/models"
	"gorm.io/gorm"
)

type saleOrderItemRepo struct {
	db *gorm.DB
}

func SaleOrderItemDBInstance(d *gorm.DB) domain.ISaleOrderItemRepository {
	return &saleOrderItemRepo{
		db: d,
	}
}

func (repo *saleOrderItemRepo) CreateSaleOrderItem(saleOrderItem *models.SaleOrderItem) error {
	if err := repo.db.Create(saleOrderItem).Error; err != nil {
		return err
	}
	return nil
}

func (repo *saleOrderItemRepo) GetSaleOrderItems(saleOrderItemID uint) []models.SaleOrderItem {
	var saleOrderItems []models.SaleOrderItem
	var err error

	if saleOrderItemID != 0 {
		err = repo.db.Where("id = ?", saleOrderItemID).Find(&saleOrderItems).Error
	} else {
		err = repo.db.Find(&saleOrderItems).Error
	}

	if err != nil {
		return []models.SaleOrderItem{}
	}
	return saleOrderItems
}

func (repo *saleOrderItemRepo) UpdateSaleOrderItem(saleOrderItem *models.SaleOrderItem) error {
	if err := repo.db.Save(saleOrderItem).Error; err != nil {
		return err
	}
	return nil
}

func (repo *saleOrderItemRepo) DeleteSaleOrderItem(saleOrderItemID uint) error {
	var saleOrderItem models.SaleOrderItem
	if err := repo.db.
		Where("id = ?", saleOrderItemID).Delete(&saleOrderItem).
		Error; err != nil {
		return err
	}
	return nil
}
