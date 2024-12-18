package models

import "time"

type Package struct {
	ID           uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name         string    `gorm:"not null"`
	SaleOrderID  uint      `gorm:"not null"`
	ShippingDate time.Time `gorm:"not null"`
	Status       string    `gorm:"not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
