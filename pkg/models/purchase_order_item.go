package models

import "time"

type PurchaseOrderItem struct {
	ID              uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	PurchaseOrderID uint      `gorm:"not null"`
	ItemID          uint      `gorm:"not null"`
	Quantity        int       `gorm:"not null"`
	CreatedAt       time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
