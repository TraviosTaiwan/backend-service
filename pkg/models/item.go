package models

import "time"

type Item struct {
	ID           uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name         string    `gorm:"not null;unique"`
	CategoryID   uint      `gorm:"not null"`
	VendorID     uint      `gorm:"not null"`
	Description  string    `gorm:"not null"`
	Quantity     int       `gorm:"not null"`
	ImageUrl     string    `gorm:"not null"`
	BuyingPrice  int       `gorm:"not null"`
	SellingPrice int       `gorm:"not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	Category     Category  `gorm:"foreignKey:CategoryID"`
	Vendor       Vendor    `gorm:"foreignKey:VendorID"`
	Tag          string    
}
