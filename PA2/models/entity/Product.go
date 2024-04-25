package entity

import "time"

type Product struct {
	Id          uint      `json:"id,omitempty"`
	Name        string    `json:"name" gorm:"type:varchar(255);uniqueIndex" validate:"required"`
	Description string    `json:"description" gorm:"type:text" validate:"required"`
	Price       float64   `json:"price" gorm:"type:float" validate:"required"`
	Image       string    `json:"image" gorm:"type:varchar(200)" validate:"required"`
	CategoryId  uint      `json:"category_id" gorm:"index"`
	Category    Category  `json:"-" gorm:"foreignKey:CategoryId"`
	AdminID     uint      `json:"admin_id" gorm:"index"`
	Admin       Admin     `json:"-" gorm:"foreignKey:AdminID"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime" db:"updated_at"`
}
