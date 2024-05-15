package models

import "time"

type BookingQueue struct {
	ID         uint      `json:"id,omitempty"`
	ProductID  uint      `gorm:"" json:"product_id"`
	Product    Product   `gorm:"foreignKey:ProductID" json:"product"`
	Count      string    `json:"total,omitempty"`
	CustomerID uint      `gorm:"" json:"customer_id"`
	Customer   Customer  `gorm:"foreignKey:CustomerID" json:"customer"`
	Taking     int       `gorm:"type:int(4)" json:"taking"`
	Status     int       `gorm:"type:int(4)" json:"status"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
