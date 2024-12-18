package domain

import (
	"github.com/iwachan14736/travios-backend-service/pkg/models"
	"github.com/iwachan14736/travios-backend-service/pkg/types"
)

type IItemRepository interface {
	GetItems(itemID uint) []models.Item
	CreateItem(item *models.Item) error
	UpdateItem(item *models.Item) error
	DeleteItem(itemID uint) error
}

type IItemService interface {
	GetItems(itemID uint) ([]types.Item, error)
	CreateItem(item *models.Item) error
	UpdateItem(item *models.Item) error
	DeleteItem(itemID uint) error
}
