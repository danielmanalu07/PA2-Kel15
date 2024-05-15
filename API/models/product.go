package models

import "time"

type Product struct {
	Id          uint      `json:"id"`
	Name        string    `json:"name" gorm:"type:varchar(225)"`
	Description string    `json:"description" gorm:"type:text"`
	Price       string    `json:"price" gorm:"type:varchar(225)"`
	Image       string    `json:"image"`
	CategoryID  uint      `json:"category_id"`
	Category    Category  `gorm:"foreignKey:CategoryID" json:"category"`
	AdminID     uint      `gorm:"" json:"admin_id"`
	Admin       Admin     `gorm:"foreignKey:AdminID" json:"admin"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
