package models

import "time"

type Table struct {
	Id        uint      `json:"id"`
	Number    int64     `json:"number" gorm:"unique"`
	Capacity  int64     `json:"capacity"`
	Image     string    `json:"image"`
	AdminID   uint      `gorm:"" json:"admin_id"`
	Admin     Admin     `gorm:"foreignKey:AdminID" json:"admin"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
