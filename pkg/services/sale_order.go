package services

import (
	"errors"

	"github.com/iwachan14736/travios-backend-service/pkg/domain"
	"github.com/iwachan14736/travios-backend-service/pkg/models"
	"github.com/iwachan14736/travios-backend-service/pkg/types"
)

type SaleOrderService struct {
	repo domain.ISaleOrderRepository
}

func SaleOrderServiceInstance(saleOrderRepo domain.ISaleOrderRepository) domain.ISaleOrderService {
	return &SaleOrderService{
		repo: saleOrderRepo,
	}
}

func (service *SaleOrderService) GetSaleOrders(saleOrderID uint) ([]types.SaleOrder, error) {
	var allSaleOrders []types.SaleOrder
	saleOrders := service.repo.GetSaleOrders(saleOrderID)
	if len(saleOrders) == 0 {
		return nil, errors.New("SaleOrder not found")
	}
	for _, val := range saleOrders {
		allSaleOrders = append(allSaleOrders, types.SaleOrder{
			ID:         val.ID,
			CustomerID: val.CustomerID,
		})
	}
	return allSaleOrders, nil
}

func (service *SaleOrderService) CreateSaleOrder(saleOrderRequest *models.SaleOrder) error {
	return service.repo.CreateSaleOrder(saleOrderRequest)
}

func (service *SaleOrderService) UpdateSaleOrder(saleOrderRequest *models.SaleOrder) error {
	return service.repo.UpdateSaleOrder(saleOrderRequest)
}

func (service *SaleOrderService) DeleteSaleOrder(saleOrderID uint) error {
	return service.repo.DeleteSaleOrder(saleOrderID)
}
