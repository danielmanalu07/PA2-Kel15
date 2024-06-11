package response

import "time"

type ProductResponse struct {
	Id          uint      `json:"id"`
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description" validate:"required"`
	Image       string    `json:"image" validate:"required"`
	Price       string    `json:"price" validate:"required"`
	CategoryID  uint      `json:"category_id"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime" db:"updated_at"`
	AdminID     uint      `json:"admin_id" gorm:"index"`
}
