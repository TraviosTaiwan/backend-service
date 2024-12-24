package domain

import (
	"github.com/iwachan14736/travios-backend-service/pkg/models"
	"github.com/iwachan14736/travios-backend-service/pkg/types"
)

type ISaleOrderRepository interface {
	GetSaleOrders(saleOrderID uint) []models.SaleOrder
	CreateSaleOrder(saleOrder *models.SaleOrder) error
	UpdateSaleOrder(saleOrder *models.SaleOrder) error
	DeleteSaleOrder(saleOrderID uint) error
}

type ISaleOrderService interface {
	GetSaleOrders(saleOrderID uint) ([]types.SaleOrder, error)
	CreateSaleOrder(saleOrder *models.SaleOrder) error
	UpdateSaleOrder(saleOrder *models.SaleOrder) error
	DeleteSaleOrder(saleOrderID uint) error
}
