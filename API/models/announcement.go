package models

import "time"

type Announcement struct {
	ID         uint      `json:"id,omitempty"`
	Title      string    `json:"title,omitempty" gorm:"type:varchar(255)"`
	Content    string    `json:"text,omitempty" gorm:"type:text"`
	AdminID    uint      `gorm:"" json:"admin_id"`
	Admin      Admin     `gorm:"foreignKey:AdminID" json:"admin"`
	Created_at time.Time `json:"created_at,omitempty" gorm:"autoCreateTime"`
	Updated_at time.Time `json:"updated_at,omitempty" gorm:"autoUpdateTime"`
}
