package repositories

import (
	"github.com/iwachan14736/travios-backend-service/pkg/domain"
	"github.com/iwachan14736/travios-backend-service/pkg/models"
	"gorm.io/gorm"
)

type saleOrderRepo struct {
	db *gorm.DB
}

func SaleOrderDBInstance(d *gorm.DB) domain.ISaleOrderRepository {
	return &saleOrderRepo{
		db: d,
	}
}

func (repo *saleOrderRepo) CreateSaleOrder(saleOrder *models.SaleOrder) error {
	if err := repo.db.Create(saleOrder).Error; err != nil {
		return err
	}
	return nil
}

func (repo *saleOrderRepo) GetSaleOrders(saleOrderID uint) []models.SaleOrder {
	var saleOrders []models.SaleOrder
	var err error

	if saleOrderID != 0 {
		err = repo.db.Where("id = ?", saleOrderID).Find(&saleOrders).Error
	} else {
		err = repo.db.Find(&saleOrders).Error
	}

	if err != nil {
		return []models.SaleOrder{}
	}
	return saleOrders
}

func (repo *saleOrderRepo) UpdateSaleOrder(saleOrder *models.SaleOrder) error {
	if err := repo.db.Save(saleOrder).Error; err != nil {
		return err
	}
	return nil
}

func (repo *saleOrderRepo) DeleteSaleOrder(saleOrderID uint) error {
	var saleOrder models.SaleOrder
	if err := repo.db.
		Where("id = ?", saleOrderID).Delete(&saleOrder).
		Error; err != nil {
		return err
	}
	return nil
}
