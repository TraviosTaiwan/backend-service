package services

import (
	"errors"

	"github.com/iwachan14736/travios-backend-service/pkg/domain"
	"github.com/iwachan14736/travios-backend-service/pkg/models"
	"github.com/iwachan14736/travios-backend-service/pkg/types"
)

type SaleOrderItemService struct {
	repo domain.ISaleOrderItemRepository
}

func SaleOrderItemServiceInstance(saleOrderItemRepo domain.ISaleOrderItemRepository) domain.ISaleOrderItemService {
	return &SaleOrderItemService{
		repo: saleOrderItemRepo,
	}
}

func (service *SaleOrderItemService) GetSaleOrderItems(saleOrderItemID uint) ([]types.SaleOrderItem, error) {
	var allSaleOrderItems []types.SaleOrderItem
	saleOrderItems := service.repo.GetSaleOrderItems(saleOrderItemID)
	if len(saleOrderItems) == 0 {
		return nil, errors.New("SaleOrderItem not found")
	}
	for _, val := range saleOrderItems {
		allSaleOrderItems = append(allSaleOrderItems, types.SaleOrderItem{
			ID:          val.ID,
			SaleOrderID: val.SaleOrderID,
			ItemID:      val.ItemID,
			Quantity:    val.Quantity,
		})
	}
	return allSaleOrderItems, nil
}

func (service *SaleOrderItemService) CreateSaleOrderItem(saleOrderItemRequest *models.SaleOrderItem) error {
	return service.repo.CreateSaleOrderItem(saleOrderItemRequest)
}

func (service *SaleOrderItemService) UpdateSaleOrderItem(saleOrderItemRequest *models.SaleOrderItem) error {
	return service.repo.UpdateSaleOrderItem(saleOrderItemRequest)
}

func (service *SaleOrderItemService) DeleteSaleOrderItem(saleOrderItemID uint) error {
	return service.repo.DeleteSaleOrderItem(saleOrderItemID)
}
