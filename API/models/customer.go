package models

import "time"

type Customer struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username" gorm:"unique"`
	Phone     string    `json:"phone" gorm:"unique"`
	Photo     string    `json:"photo" gorm:"not null"`
	Address   string    `json:"address" gorm:"not null"`
	Password  string    `json:"password" gorm:"min:6|confirmed"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
