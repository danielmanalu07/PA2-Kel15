package entity

import "time"

type Cart struct {
	Id         uint      `json:"id"`
	CustomerID uint      `json:"customer_id" gorm:"index"`
	Customer   Customer  `json:"-" gorm:"foreignKey:customer_id"`
	ProductID  uint      `json:"product_id" gorm:"index"`
	Product    Product   `json:"product" gorm:"foreignKey:product_id"`
	Quantity   int       `json:"quantity" gorm:"default:1"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime" db:"updated_at"`
}
