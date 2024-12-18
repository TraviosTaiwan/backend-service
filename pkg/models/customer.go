package models

import "time"

type Customer struct {
	ID        uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Phone     string `gorm:"unique"`
	Platform  string `gorm:"not null"`
	Username  string `gorm:"not null"`
	Name      string `gorm:"not null"`
	Address   string
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
