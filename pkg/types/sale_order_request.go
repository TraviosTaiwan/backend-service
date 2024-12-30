package types

import validation "github.com/go-ozzo/ozzo-validation/v4"

type CartItem struct {
	ItemID   uint `json:"item_id"`
	Quantity int  `json:"quantity"`
}

type SaleOrderRequest struct {
	CustomerName    string     `json:"customer_name"`
	CustomerPhone   string     `json:"customer_phone"`
	CustomerAddress string     `json:"customer_address"`
	CartItems       []CartItem `json:"cart_items"`
	Remark          string     `json:"customer_remark"`
}

func (req SaleOrderRequest) ValidateSaleOrderRequest() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.CustomerName,
			validation.Required.Error("Customer name cannot be empty"),
		),
		validation.Field(&req.CustomerPhone,
			validation.Required.Error("Customer phone cannot be empty"),
		),
		validation.Field(&req.CartItems,
			validation.Required.Error("Cart items cannot be empty"),
		),
	)
} 