package models

import "time"

type Customer struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Phone     string    `gorm:"unique;not null" json:"phone"`
	Platform  string    `gorm:"not null" json:"platform"`
	Username  string    `gorm:"not null" json:"username"`
	Name      string    `gorm:"not null" json:"name"`
	Address   string    `json:"address"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
