package services

import (
	"errors"

	"github.com/iwachan14736/travios-backend-service/pkg/domain"
	"github.com/iwachan14736/travios-backend-service/pkg/models"
	"github.com/iwachan14736/travios-backend-service/pkg/types"
)

type ItemService struct {
	repo domain.IItemRepository
}

func ItemServiceInstance(itemRepo domain.IItemRepository) domain.IItemService {
	return &ItemService{
		repo: itemRepo,
	}
}

func (service *ItemService) GetItems(itemID uint) ([]types.Item, error) {
	var allItems []types.Item
	items := service.repo.GetItems(itemID)
	if len(items) == 0 {
		return nil, errors.New("Item not found")
	}
	for _, val := range items {
		allItems = append(allItems, types.Item{
			ID:           val.ID,
			Name:         val.Name,
			CategoryID:   val.CategoryID,
			VendorID:     val.VendorID,
			Quantity:     val.Quantity,
			BuyingPrice:  val.BuyingPrice,
			SellingPrice: val.SellingPrice,
			Description:  val.Description,
			ImageUrl:     val.ImageUrl,
			Tag:          val.Tag,
		})
	}
	return allItems, nil
}

func (service *ItemService) GetItemByTag(tag string) ([]types.Item, error) {
	var allItems []types.Item
	items, err := service.repo.GetItemByTag(tag)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, errors.New("Items not found")
	}
	for _, val := range items {
		allItems = append(allItems, types.Item{
			ID:           val.ID,
			Name:         val.Name,
			CategoryID:   val.CategoryID,
			VendorID:     val.VendorID,
			Quantity:     val.Quantity,
			BuyingPrice:  val.BuyingPrice,
			SellingPrice: val.SellingPrice,
			Description:  val.Description,
			ImageUrl:     val.ImageUrl,
			Tag:          val.Tag,
		})
	}
	return allItems, nil
}

func (service *ItemService) CreateItem(itemRequest *models.Item) error {
	return service.repo.CreateItem(itemRequest)
}

func (service *ItemService) UpdateItem(itemRequest *models.Item) error {
	return service.repo.UpdateItem(itemRequest)
}

func (service *ItemService) DeleteItem(itemID uint) error {
	return service.repo.DeleteItem(itemID)
}
