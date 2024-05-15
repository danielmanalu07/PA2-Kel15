package models

import "time"

type Category struct {
	Id          uint      `json:"id"`
	Name        string    `json:"name" gorm:"type:varchar(225)"`
	Description string    `json:"description" gorm:"type:text"`
	AdminID     uint      `gorm:"" json:"admin_id"`
	Admin       Admin     `gorm:"foreignKey:AdminID" json:"admin"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
