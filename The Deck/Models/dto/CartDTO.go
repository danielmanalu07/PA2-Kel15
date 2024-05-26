package dto

import "time"

type RequestCartCreate struct {
	Id        uint      `json:"id"`
	ProductID uint      `json:"product_id" gorm:"index"`
	Quantity  int       `json:"quantity" gorm:"default:1"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime" db:"updated_at"`
}

type RequestCartUpdate struct {
	Quantity int `json:"quantity" gorm:"default:1"`
}
