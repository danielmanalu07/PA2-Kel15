package entity

import "time"

type Product struct {
	Id          uint      `json:"id"`
	Name        string    `json:"name" gorm:"type:varchar(50);index:idx_nm,unique" validate:"required"`
	Description string    `json:"description" gorm:"type:text" validate:"required"`
	Image       string    `json:"image" gorm:"type:varchar(50)" validate:"required"`
	Price       string    `json:"price" gorm:"type:varchar(50)" validate:"required"`
	CategoryID  uint      `json:"category_id" gorm:"index"`
	Category    Category  `json:"-" gorm:"foreignKey:category_id"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime" db:"updated_at"`
	AdminID     uint      `json:"admin_id" gorm:"index"`
	Admin       Admin     `json:"-" gorm:"foreignKey:admin_id"`
}
