package models

import "time"

type PurchaseOrder struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	VendorID   uint      `gorm:"not null" json:"vendor_id"`
	TotalPrice int       `gorm:"not null" json:"total_price"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
