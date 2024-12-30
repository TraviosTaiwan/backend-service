package types

import validation "github.com/go-ozzo/ozzo-validation/v4"

type SaleOrder struct {
	ID         uint   `json:"id"`
	CustomerID uint   `json:"customer_id"`
	Remark     string `json:"remark"`
}

func (saleOrder SaleOrder) ValidateSaleOrder() error {
	return validation.ValidateStruct(&saleOrder,
		validation.Field(&saleOrder.CustomerID,
			validation.Required.Error("Customer ID cannot be empty"),
		),
	)
}
