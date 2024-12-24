package models

import "time"

type SaleOrder struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	CustomerID uint      `gorm:"not null" json:"customer_id"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
