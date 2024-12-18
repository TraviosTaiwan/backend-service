package models

import "time"

type SaleOrderItem struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	SaleOrderID uint      `gorm:"not null" json:"sale_order_id"`
	ItemID      uint      `gorm:"not null" json:"item_id"`
	Quantity    int       `gorm:"not null" json:"quantity"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
