package domain

import (
	"github.com/iwachan14736/travios-backend-service/pkg/models"
	"github.com/iwachan14736/travios-backend-service/pkg/types"
)

type ISaleOrderItemRepository interface {
	GetSaleOrderItems(saleOrderItemID uint) []models.SaleOrderItem
	CreateSaleOrderItem(saleOrderItem *models.SaleOrderItem) error
	UpdateSaleOrderItem(saleOrderItem *models.SaleOrderItem) error
	DeleteSaleOrderItem(saleOrderItemID uint) error
}

type ISaleOrderItemService interface {
	GetSaleOrderItems(saleOrderItemID uint) ([]types.SaleOrderItem, error)
	CreateSaleOrderItem(saleOrderItem *models.SaleOrderItem) error
	UpdateSaleOrderItem(saleOrderItem *models.SaleOrderItem) error
	DeleteSaleOrderItem(saleOrderItemID uint) error
}
