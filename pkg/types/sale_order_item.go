package types

import validation "github.com/go-ozzo/ozzo-validation/v4"

type SaleOrderItem struct {
	ID          uint `json:"id"`
	SaleOrderID uint `json:"sale_order_id"`
	ItemID      uint `json:"item_id"`
	Quantity    int  `json:"quantity"`
}

func (saleOrderItem SaleOrderItem) ValidateSaleOrderItem() error {
	return validation.ValidateStruct(&saleOrderItem,
		validation.Field(&saleOrderItem.SaleOrderID,
			validation.Required.Error("Sale Order ID cannot be empty"),
		),
		validation.Field(&saleOrderItem.ItemID,
			validation.Required.Error("Item ID cannot be empty"),
		),
		validation.Field(&saleOrderItem.Quantity,
			validation.Required.Error("Quantity cannot be empty"),
		),
	)
}
