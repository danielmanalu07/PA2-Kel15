package entity

import "time"

type Category struct {
	Id          uint      `json:"id"`
	Name        string    `json:"name" gorm:"type:varchar(255);uniqueIndex" validate:"required"`
	Description string    `json:"description" gorm:"type:text" validate:"required"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime" db:"updated_at"`
	AdminID     uint      `json:"admin_id" gorm:"index"`
	Admin       Admin     `json:"-" gorm:"foreignKey:admin_id"`
}
