package types

import validation "github.com/go-ozzo/ozzo-validation/v4"

type Item struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	CategoryID   uint   `json:"category_id"`
	VendorID     uint   `json:"vendor_id"`
	Description  string `json:"description"`
	Quantity     int    `json:"quantity"`
	ImageUrl     string `json:"image_url"`
	BuyingPrice  int    `json:"buying_price"`
	SellingPrice int    `json:"selling_price"`
	Tag          string `json:"tag"`
}

func (item Item) ValidateItem() error {
	return validation.ValidateStruct(&item,
		validation.Field(&item.Name,
			validation.Required.Error("Name cannot be empty"),
			validation.Length(1, 50),
		),
		validation.Field(&item.CategoryID,
			validation.Required.Error("Category ID cannot be empty"),
		),
		validation.Field(&item.VendorID,
			validation.Required.Error("Vendor ID cannot be empty"),
		),
		validation.Field(&item.SellingPrice,
			validation.Required.Error("Selling price cannot be empty"),
		),
		validation.Field(&item.Quantity,
			validation.Required.Error("Quantity cannot be empty"),
		),
	)
}
